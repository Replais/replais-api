# replais-api

🔥 What the product is

ReplAIs is a browser extension + backend service that helps users reply across any platform:
	•	WhatsApp
	•	Gmail
	•	Instagram
	•	Slack
	•	LinkedIn
	•	Twitter DMs
	•	…and more

It understands:
	•	The conversation context (messages, participants)
	•	The platform (WhatsApp, Gmail, etc.)
	•	The relationship (friend, boss, crush)
	•	The user’s personal tone/persona
	•	The conversation’s vibe and constraints

Then generates the perfect reply, on-tone and on-context, instantly.

Think of it as “Grammarly for Replies”.

⸻

🎯 What the replais-api backend is supposed to do

The backend is the brain + database behind the extension.

It must:

1. Accept a unified incoming request from the extension

This includes:
	•	Conversation context
	•	User’s persona selection
	•	Tone setting
	•	Past contact instructions
	•	Reply-specific instructions
	•	UI state (draft, selection, etc.)

The FE sends a structured JSON payload called ReplaisRequest.

⸻

2. Understand/store per-contact settings

Examples:
	•	“With Shivam on WhatsApp, my tone = casual & persona = buddy”
	•	“With Boss on Slack, tone = formal”
	•	“With Crush on Instagram, persona = flirty, tone = playful”

This data lives in the backend, not extension.

⸻

3. Build a prompt & call the AI model

Backend merges:
	•	Persona prompt
	•	Tone constraints
	•	Past instructions
	•	Context (messages)
	•	Reply-only instructions

For now this can be stubbed.

⸻

4. Return a smart reply

replyText: string

And optionally:

updatedContactConfig
(if user changed persona/tone/instructions)

⸻

5. Save request + response logs

For analytics, debugging, billing, LLM cost calculation.

⸻

🧠 High-level purpose of this backend

The backend acts as:
	1.	State store for users + contacts + settings
	2.	Prompt builder
	3.	LLM request router
	4.	Request logger
	5.	Future: rate limits, auth, billing, analytics

The browser extension becomes dumb.
The backend becomes the intelligent center.

Backend style:
	•	Clean Go code
	•	Chi router
	•	PostgreSQL
	•	Migrations
	•	DTOs
	•	No complicated auth yet
	•	No queues for now