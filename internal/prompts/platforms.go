package prompts

type Platform string

const (
	PlatformWhatsApp  Platform = "whatsapp"
	PlatformGmail     Platform = "gmail"
	PlatformLinkedIn  Platform = "linkedin"
	PlatformInstagram Platform = "instagram"
	PlatformBumble    Platform = "bumble"
	PlatformSlack     Platform = "slack"
	PlatformTwitter   Platform = "twitter"
)

// PlatformPrompts contains system-level prompts for each platform.
// These prompts are injected into the AI prompt based on the platform (WhatsApp, Gmail, etc.)
// to guide the AI on platform-specific messaging conventions and styles.
var PlatformPrompts = map[Platform]string{
	PlatformWhatsApp: `You are replying on WhatsApp. Keep messages concise and casual. Use emojis naturally when appropriate. WhatsApp messages are typically short, conversational, and informal. Avoid long paragraphs unless necessary.`,

	PlatformGmail: `You are replying to an email on Gmail. Emails can be more formal and structured than instant messages. Use proper greetings and sign-offs. Structure longer emails with paragraphs. Be professional but match the tone of the original email.`,

	PlatformLinkedIn: `You are replying on LinkedIn. Maintain a professional tone. LinkedIn is a business networking platform, so be respectful, clear, and professional. Avoid overly casual language unless the conversation context clearly calls for it.`,

	PlatformInstagram: `You are replying on Instagram. Instagram DMs are typically casual and friendly. Use emojis naturally. Keep messages relatively short and engaging. Match the visual, creative nature of the platform.`,

	PlatformBumble: `You are replying on Bumble. Keep messages light, engaging, and friendly. Show interest without being too intense. Ask questions to keep the conversation flowing. Be authentic and avoid generic responses.`,

	PlatformSlack: `You are replying on Slack. Slack messages can range from casual to professional depending on the workspace and channel. Be concise and clear. Use threads for longer discussions. Match the tone of your workplace culture.`,

	PlatformTwitter: `You are replying to a Twitter DM. Twitter DMs are typically casual and brief. Keep messages concise and engaging. Match the conversational, public-facing nature of Twitter.`,
}

func GetPlatformPrompt(platform Platform) string {
	if prompt, exists := PlatformPrompts[platform]; exists {
		return prompt
	}
	return ""
}
