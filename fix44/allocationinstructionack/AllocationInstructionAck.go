//Package allocationinstructionack msg type = P.
package allocationinstructionack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a AllocationInstructionAck wrapper for the generic Message type
type Message struct {
	message.Message
}

//AllocID is a required field for AllocationInstructionAck.
func (m Message) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationInstructionAck.
func (m Message) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationInstructionAck.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationInstructionAck.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationInstructionAck.
func (m Message) SecondaryAllocID() (*field.SecondaryAllocIDField, errors.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationInstructionAck.
func (m Message) GetSecondaryAllocID(f *field.SecondaryAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for AllocationInstructionAck.
func (m Message) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationInstructionAck.
func (m Message) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for AllocationInstructionAck.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationInstructionAck.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationInstructionAck.
func (m Message) AllocStatus() (*field.AllocStatusField, errors.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationInstructionAck.
func (m Message) GetAllocStatus(f *field.AllocStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationInstructionAck.
func (m Message) AllocRejCode() (*field.AllocRejCodeField, errors.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationInstructionAck.
func (m Message) GetAllocRejCode(f *field.AllocRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a non-required field for AllocationInstructionAck.
func (m Message) AllocType() (*field.AllocTypeField, errors.MessageRejectError) {
	f := &field.AllocTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from AllocationInstructionAck.
func (m Message) GetAllocType(f *field.AllocTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
func (m Message) AllocIntermedReqType() (*field.AllocIntermedReqTypeField, errors.MessageRejectError) {
	f := &field.AllocIntermedReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationInstructionAck.
func (m Message) GetAllocIntermedReqType(f *field.AllocIntermedReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for AllocationInstructionAck.
func (m Message) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from AllocationInstructionAck.
func (m Message) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationInstructionAck.
func (m Message) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationInstructionAck.
func (m Message) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationInstructionAck.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationInstructionAck.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationInstructionAck.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationInstructionAck.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationInstructionAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationInstructionAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationInstructionAck.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationInstructionAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationInstructionAck.
func (m Message) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationInstructionAck.
func (m Message) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AllocationInstructionAck messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AllocationInstructionAck.
func Builder(
	allocid *field.AllocIDField,
	transacttime *field.TransactTimeField,
	allocstatus *field.AllocStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("P"))
	builder.Body().Set(allocid)
	builder.Body().Set(transacttime)
	builder.Body().Set(allocstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "P", r
}