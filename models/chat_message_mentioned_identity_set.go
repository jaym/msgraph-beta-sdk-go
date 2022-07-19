package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageMentionedIdentitySet 
type ChatMessageMentionedIdentitySet struct {
    IdentitySet
    // If present, represents a conversation (for example, team or channel) @mentioned in a message.
    conversation TeamworkConversationIdentityable
    // If present, represents a tag @mentioned in a team message.
    tag TeamworkTagIdentityable
}
// NewChatMessageMentionedIdentitySet instantiates a new ChatMessageMentionedIdentitySet and sets the default values.
func NewChatMessageMentionedIdentitySet()(*ChatMessageMentionedIdentitySet) {
    m := &ChatMessageMentionedIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.chatMessageMentionedIdentitySet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateChatMessageMentionedIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageMentionedIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageMentionedIdentitySet(), nil
}
// GetConversation gets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
func (m *ChatMessageMentionedIdentitySet) GetConversation()(TeamworkConversationIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.conversation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageMentionedIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["conversation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConversationIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversation(val.(TeamworkConversationIdentityable))
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkTagIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val.(TeamworkTagIdentityable))
        }
        return nil
    }
    return res
}
// GetTag gets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) GetTag()(TeamworkTagIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
// Serialize serializes information the current object
func (m *ChatMessageMentionedIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conversation", m.GetConversation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConversation sets the conversation property value. If present, represents a conversation (for example, team or channel) @mentioned in a message.
func (m *ChatMessageMentionedIdentitySet) SetConversation(value TeamworkConversationIdentityable)() {
    if m != nil {
        m.conversation = value
    }
}
// SetTag sets the tag property value. If present, represents a tag @mentioned in a team message.
func (m *ChatMessageMentionedIdentitySet) SetTag(value TeamworkTagIdentityable)() {
    if m != nil {
        m.tag = value
    }
}
