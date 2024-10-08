package session

import (
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/text"
)

func (b *Broadcast) SecureChatMessage(session Session, pk play.ChatMessage, index int32) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	b.prev_msgs_mu.Lock()
	defer b.prev_msgs_mu.Unlock()

	for _, ses := range b.sessions {
		ses.PlayerChatMessage(pk, session, "minecraft:chat", index, b.previousMessages)
	}
	b.appendMessage()
}

func (b *Broadcast) appendMessage() {
	if len(b.previousMessages) != 20 {
		b.previousMessages = append([]play.PreviousMessage{{MessageID: int32(len(b.previousMessages))}}, b.previousMessages...)
	}
}

func (b *Broadcast) DisguisedChatMessage(session Session, content text.TextComponent) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, ses := range b.sessions {
		ses.DisguisedChatMessage(content, session, "minecraft:chat")
	}
}
