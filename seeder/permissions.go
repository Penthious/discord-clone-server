package seeder

import (
	"discord-clone-server/models"

	"gorm.io/gorm"
)

func PermissionsSeeder(db *gorm.DB) error {
	permissions := []models.Permission{
		{
			Name:       "Admin",
			Permission: "admin",
			Detail:     "Members with this permission have every permission and also bypass channel specific permissions. This is a dangerous permission to grant.",
			Type:       "general",
		},
		{
			Name:       "View Audit Log",
			Permission: "view_audit_log",
			Detail:     "Members with this permission have access to view the server's audit logs.",
			Type:       "general",
		},
		{
			Name:       "Manage Server",
			Permission: "manage_server",
			Detail:     "Members with this permission can change the server's name or move regions.",
			Type:       "general",
		},
		{
			Name:       "Manage Roles",
			Permission: "manage_roles",
			Detail:     "Members with this permission can create new roles and edit/delete roles lower than this one.",
			Type:       "general",
		},
		{
			Name:       "Manage Channels",
			Permission: "manage_channels",
			Detail:     "Members with this permission can create new channels and edit or delete existing ones.",
			Type:       "general",
		},
		{
			Name:       "Kick Members",
			Permission: "kick_members",
			Detail:     "Members with this permission can kick members from the server.",
			Type:       "general",
		},
		{
			Name:       "Ban Members",
			Permission: "ban_members",
			Detail:     "Members with this permission can ban members from the server.",
			Type:       "general",
		},
		{
			Name:       "Create Invite",
			Permission: "create_invite",
			Detail:     "Members with this permission can invite other users",
			Type:       "general",
		},
		{
			Name:       "Change Nickname",
			Permission: "change_nickname",
			Detail:     "Members with this permission can change users nicknames",
			Type:       "general",
		},
		{
			Name:       "Manage Emojis",
			Permission: "manage_emojis",
			Detail:     "Members with this permission can add/update/delete emojis",
			Type:       "general",
		},
		{
			Name:       "Manage Webhooks",
			Permission: "manage_webhooks",
			Detail:     "Members with this permission can add/update/delete webhooks",
			Type:       "general",
		},
		{
			Name:       "Read Text Channels & See Voice Channels",
			Permission: "read_channels",
			Detail:     "Members with this permission can access public text/voice channels",
			Type:       "general",
		},
		{
			Name:       "Send Messages",
			Permission: "send_message",
			Detail:     "Members with this permission can send messages to other users",
			Type:       "text",
		},
		{
			Name:       "Send TTS Messages",
			Permission: "send_tts_message",
			Detail:     "Members with this permission can send text-to-speech mesages by starting a message with /tts. These messages can be heard by everyone focused on the channel.",
			Type:       "text",
		},
		{
			Name:       "Manage Messages",
			Permission: "manage_message",
			Detail:     "Members with this permission can delete messages by other members or pin any message.",
			Type:       "text",
		},
		{
			Name:       "Embed Links",
			Permission: "embed_link",
			Detail:     "Members with this permission can embed links",
			Type:       "text",
		},
		{
			Name:       "Attach Files",
			Permission: "attach_file",
			Detail:     "Members with this permission can upload files",
			Type:       "text",
		},
		{
			Name:       "Read Message History",
			Permission: "read_message_history",
			Detail:     "Members with this permission can load previous messages",
			Type:       "text",
		},
		{
			Name:       "Mention @everyone, @here, and All Roles",
			Permission: "use_mentions",
			Detail:     "Members with this permission can use @everyone or @here to ping all members in this channel. They can also @mention all roles, even if the role's (Allow anyone to mention this role) permission is disabled",
			Type:       "text",
		},
		{
			Name:       "Use External Emojis",
			Permission: "external_emojis",
			Detail:     "Members with this permission can use emojis from other servers in this server",
			Type:       "text",
		},
		{
			Name:       "Add Reactions",
			Permission: "add_reaction",
			Detail:     "Members with this permission can add new reactions to a message. Members can still react using reactions already added to messages without this permission.",
			Type:       "text",
		},
		{
			Name:       "Connect",
			Permission: "connect",
			Detail:     "Members with this permission can connect to voice calls",
			Type:       "voice",
		},
		{
			Name:       "Speak",
			Permission: "speak",
			Detail:     "Members with this permission can speak in voice calls",
			Type:       "voice",
		},
		{
			Name:       "Video",
			Permission: "video",
			Detail:     "Members with this permission can stream into this server",
			Type:       "voice",
		},
		{
			Name:       "Mute Members",
			Permission: "mute",
			Detail:     "Members with this permission can mute other members",
			Type:       "voice",
		},
		{
			Name:       "Deafen Members",
			Permission: "deafen",
			Detail:     "Members with this permission can lower other members volume",
			Type:       "voice",
		},
		{
			Name:       "Move Members",
			Permission: "move",
			Detail:     "Members with this permission can drag other members out of this channel. They can only move members between channels both they and the member they are moving have access.",
			Type:       "voice",
		},
		{
			Name:       "Use Voice Activity",
			Permission: "push_to_talk",
			Detail:     "Members must use Push-to-talk in this channel if this permission is active",
			Type:       "voice",
		},
		{
			Name:       "Priority Speaker",
			Permission: "priority_speaker",
			Detail:     "Users with this permission have the ability to be more easily heard when talking. When activated the volume of others without this permission will be automatically lowered. Priority Speaker is activated by using the ",
			Type:       "voice",
		},
	}

	return db.Create(permissions).Error
}
