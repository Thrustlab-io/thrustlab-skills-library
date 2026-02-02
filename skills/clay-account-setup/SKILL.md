---
name: clay-account-setup
description: Guides client through creating a Clay.com account using Thrustlab referral link and adding team members with proper permissions
argument-hint: "client name, client email"
---

# Clay.com Account Setup Skill

You are helping set up a new Clay.com account for a GTM client using Thrustlab's referral link and adding your team as admins for ongoing support.

## Required Information

Gather from the user:
1. **Client name** - The name of the client/company
2. **Client email** - The email address that will own the Clay workspace
3. **Thrustlab team emails** - Who from your team needs admin access

## Step 1: Create Clay Account with Referral

Provide these instructions to the client:

```
📋 Clay.com Account Creation

1. Visit Clay.com via Thrustlab's referral link:
   🔗 https://clay.com/?via=thrustlab

2. Click "Sign Up" or "Get Started"

3. Sign up with YOUR COMPANY EMAIL: [Insert Client Email]
   ⚠️ Important: Use your company email so you own the workspace

4. Complete the signup form:
   - Full Name: [Client Name]
   - Company: [Client Company]
   - Role: [Client Role]
   - Use Case: Select "Outbound Sales" or "Lead Generation"

5. Verify your email address (check inbox for verification link)

6. Choose your plan:

   💡 Plan Recommendations:

   • Starter ($149/mo)
     - 500 credits/month
     - 1 workspace
     - Best for: Testing Clay or small teams (<5 people)

   • Explorer ($349/mo) ⭐ MOST POPULAR
     - 2,000 credits/month
     - Unlimited users
     - Best for: Growing GTM teams with regular outbound

   • Pro ($800/mo)
     - 12,000 credits/month
     - Unlimited users + advanced features
     - Best for: Scaled operations with high-volume enrichment

7. Complete payment setup

8. 🎉 You'll receive 500 FREE Clay credits as a welcome bonus!
```

## Step 2: Initial Workspace Configuration

After account creation, guide them through basic setup:

```
📋 Workspace Configuration

1. You'll land on the Clay dashboard

2. Set your workspace name:
   - Click workspace name (top left, probably says "My Workspace")
   - Rename to: "[Client Company] - GTM Operations"
   - Click "Save"

3. Complete your profile:
   - Click your avatar (top right)
   - Go to "Settings" → "Profile"
   - Add profile photo (optional but recommended)
   - Confirm email and name
```

## Step 3: Add Thrustlab Team Members

Critical step - guide them to add your team as admins:

```
📋 Add Thrustlab Team as Admins

1. In Clay, click your workspace name (top left)

2. Select "Settings" from the dropdown

3. Go to "Members" tab

4. Click "Invite Members"

5. Add Thrustlab team members:

   [LIST YOUR TEAM EMAILS HERE - REPLACE PLACEHOLDERS]

   👤 [TEAM MEMBER EMAIL 1]
   - Role: Admin
   - Reason: Primary support and strategy

   👤 [TEAM MEMBER EMAIL 2]
   - Role: Admin
   - Reason: Technical setup and automation

   👤 [TEAM MEMBER EMAIL 3]
   - Role: Member
   - Reason: Collaboration and training

6. Permissions Explanation:

   🔑 Admin
   - Full workspace access
   - Can manage billing and members
   - Can create/edit/delete all tables
   - Can configure integrations
   → Use for: Thrustlab core team

   📝 Member
   - Can create and edit tables
   - Cannot manage billing or remove admins
   - Can use all Clay features
   → Use for: Thrustlab collaborators

   👁️ Viewer
   - Read-only access
   - Can view tables but not edit
   - Cannot create new tables
   → Use for: Stakeholders who just need visibility

7. Click "Send Invitations"

8. Notify Thrustlab team:
   - Send them a message that invitations are sent
   - They'll receive email invites to accept
```

## Step 4: Verify Access

Confirm everything is set up correctly:

```
✅ Verification Checklist

Ask the client to verify:

□ Clay account created with their company email
□ Email address verified (check for verification email)
□ Plan selected and payment completed
□ Workspace renamed to "[Company] - GTM Operations"
□ All Thrustlab team members invited
□ Invitations sent successfully (check for confirmation)
□ Received 500 free welcome credits

Ask Thrustlab team to verify:
□ Received invitation emails from Clay
□ Accepted invitations
□ Can access the workspace
□ Have Admin permissions (can see "Settings" and "Members")
```

## Step 5: Next Steps

After account setup is complete:

```
🎯 What's Next?

Now that your Clay account is ready, you can:

1. Set up your workspace structure:
   Run: /clay-tables-setup [Client Name]
   → Creates Target Accounts, Contacts, Outreach, and more

2. Configure automation workflows:
   Run: /clay-automation-setup [Client Name]
   → Sets up enrichment and notification automations

3. Connect integrations:
   Run: /clay-integrations-setup [Client Name]
   → Connects your CRM, email tools, and Slack

💡 Tip: Complete these in order for best results!
```

## Troubleshooting

### Issue: Referral link not applying credits
**Solution**:
- Ensure using exact link: `https://clay.com/?via=thrustlab`
- Don't click other Clay links before signing up
- Sign up in a fresh browser session or incognito mode
- Contact Clay support if credits don't appear after 24 hours

### Issue: Can't add team members
**Solution**:
- Verify your email address first
- Ensure payment is completed
- Check that you have "Admin" role (you should as workspace owner)
- Try refreshing the page

### Issue: Team member invitations not received
**Solution**:
- Check spam/junk folders
- Verify email addresses are correct
- Resend invitations from Members page
- Have team members check their email filters

### Issue: Wrong person created the account
**Solution**:
- Contact Clay support to transfer ownership
- Or: Create new account with correct email, cancel old one
- Referral credits should still apply to new account

## Expected Outcomes

By the end of this setup:
- ✅ Clay.com account created via Thrustlab referral link
- ✅ Client owns the workspace with their company email
- ✅ Appropriate plan selected based on usage needs
- ✅ 500 free credits received
- ✅ Thrustlab team added with Admin permissions
- ✅ All team members can access the workspace
- ✅ Ready for table and automation setup

## Notes

- **Workspace ownership**: Client should own the account, not Thrustlab
- **Admin access**: Thrustlab needs Admin role to configure everything properly
- **Billing**: Client is responsible for Clay subscription costs
- **Credits**: Used for enrichment operations (email finding, company data, etc.)
- **Credit consumption**: Typical usage is 2-5 credits per contact enriched
- **Referral benefit**: Both you and client get bonus credits

## After Setup

Confirm with the user:
```
✅ Clay account setup complete!

Next step: Set up your workspace tables
Run: /clay-tables-setup [Client Name]

This will create:
- Target Accounts database
- Contacts database
- Outreach campaigns table
- Competitive intelligence
- ICP analysis table
```

---

**Success Metrics:**
- Account created with referral link
- Thrustlab team has admin access
- Ready to build workspace structure
