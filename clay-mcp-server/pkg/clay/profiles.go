package clay

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// WorkspaceProfile represents a saved workspace configuration
type WorkspaceProfile struct {
	WorkspaceID     string `json:"workspace_id"`
	SessionCookie   string `json:"session_cookie"`
	FrontendVersion string `json:"frontend_version,omitempty"`
	Description     string `json:"description,omitempty"`
}

// ProfilesConfig represents the workspace profiles configuration file
type ProfilesConfig struct {
	Workspaces map[string]*WorkspaceProfile `json:"workspaces"`
	Default    string                       `json:"default,omitempty"`
}

// ProfileManager handles workspace profile operations
type ProfileManager struct {
	configPath string
	config     *ProfilesConfig
}

// NewProfileManager creates a new profile manager
func NewProfileManager() (*ProfileManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".clay")
	configPath := filepath.Join(configDir, "workspaces.json")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	pm := &ProfileManager{
		configPath: configPath,
	}

	// Load existing config or create new one
	if err := pm.load(); err != nil {
		if os.IsNotExist(err) {
			pm.config = &ProfilesConfig{
				Workspaces: make(map[string]*WorkspaceProfile),
			}
			if err := pm.save(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return pm, nil
}

// load reads the profiles configuration from disk
func (pm *ProfileManager) load() error {
	data, err := os.ReadFile(pm.configPath)
	if err != nil {
		return err
	}

	var config ProfilesConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse profiles config: %w", err)
	}

	if config.Workspaces == nil {
		config.Workspaces = make(map[string]*WorkspaceProfile)
	}

	pm.config = &config
	return nil
}

// save writes the profiles configuration to disk
func (pm *ProfileManager) save() error {
	data, err := json.MarshalIndent(pm.config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal profiles config: %w", err)
	}

	if err := os.WriteFile(pm.configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write profiles config: %w", err)
	}

	return nil
}

// AddProfile adds or updates a workspace profile
func (pm *ProfileManager) AddProfile(name string, profile *WorkspaceProfile) error {
	pm.config.Workspaces[name] = profile
	return pm.save()
}

// RemoveProfile removes a workspace profile
func (pm *ProfileManager) RemoveProfile(name string) error {
	if _, exists := pm.config.Workspaces[name]; !exists {
		return fmt.Errorf("profile '%s' does not exist", name)
	}

	delete(pm.config.Workspaces, name)

	// Clear default if it was the removed profile
	if pm.config.Default == name {
		pm.config.Default = ""
	}

	return pm.save()
}

// GetProfile retrieves a workspace profile by name
func (pm *ProfileManager) GetProfile(name string) (*WorkspaceProfile, error) {
	profile, exists := pm.config.Workspaces[name]
	if !exists {
		return nil, fmt.Errorf("profile '%s' does not exist", name)
	}
	return profile, nil
}

// ListProfiles returns all workspace profiles
func (pm *ProfileManager) ListProfiles() map[string]*WorkspaceProfile {
	return pm.config.Workspaces
}

// SetDefault sets the default workspace profile
func (pm *ProfileManager) SetDefault(name string) error {
	if _, exists := pm.config.Workspaces[name]; !exists {
		return fmt.Errorf("profile '%s' does not exist", name)
	}

	pm.config.Default = name
	return pm.save()
}

// GetDefault returns the default workspace profile name
func (pm *ProfileManager) GetDefault() string {
	return pm.config.Default
}

// GetDefaultProfile returns the default workspace profile
func (pm *ProfileManager) GetDefaultProfile() (*WorkspaceProfile, error) {
	if pm.config.Default == "" {
		return nil, fmt.Errorf("no default workspace profile set")
	}
	return pm.GetProfile(pm.config.Default)
}
