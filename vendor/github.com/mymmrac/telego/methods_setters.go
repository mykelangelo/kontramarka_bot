package telego

// WithOffset adds offset parameter
func (p *GetUpdatesParams) WithOffset(offset int) *GetUpdatesParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetUpdatesParams) WithLimit(limit int) *GetUpdatesParams {
	p.Limit = limit
	return p
}

// WithTimeout adds timeout parameter
func (p *GetUpdatesParams) WithTimeout(timeout int) *GetUpdatesParams {
	p.Timeout = timeout
	return p
}

// WithAllowedUpdates adds allowed updates parameter
func (p *GetUpdatesParams) WithAllowedUpdates(allowedUpdates ...string) *GetUpdatesParams {
	p.AllowedUpdates = allowedUpdates
	return p
}

// WithURL adds URL parameter
func (p *SetWebhookParams) WithURL(url string) *SetWebhookParams {
	p.URL = url
	return p
}

// WithCertificate adds certificate parameter
func (p *SetWebhookParams) WithCertificate(certificate *InputFile) *SetWebhookParams {
	p.Certificate = certificate
	return p
}

// WithIPAddress adds ip address parameter
func (p *SetWebhookParams) WithIPAddress(ipAddress string) *SetWebhookParams {
	p.IPAddress = ipAddress
	return p
}

// WithMaxConnections adds max connections parameter
func (p *SetWebhookParams) WithMaxConnections(maxConnections int) *SetWebhookParams {
	p.MaxConnections = maxConnections
	return p
}

// WithAllowedUpdates adds allowed updates parameter
func (p *SetWebhookParams) WithAllowedUpdates(allowedUpdates ...string) *SetWebhookParams {
	p.AllowedUpdates = allowedUpdates
	return p
}

// WithDropPendingUpdates adds drop pending updates parameter
func (p *SetWebhookParams) WithDropPendingUpdates() *SetWebhookParams {
	p.DropPendingUpdates = true
	return p
}

// WithDropPendingUpdates adds drop pending updates parameter
func (p *DeleteWebhookParams) WithDropPendingUpdates() *DeleteWebhookParams {
	p.DropPendingUpdates = true
	return p
}

// WithChatID adds chat ID parameter
func (p *SendMessageParams) WithChatID(chatID ChatID) *SendMessageParams {
	p.ChatID = chatID
	return p
}

// WithText adds text parameter
func (p *SendMessageParams) WithText(text string) *SendMessageParams {
	p.Text = text
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendMessageParams) WithParseMode(parseMode string) *SendMessageParams {
	p.ParseMode = parseMode
	return p
}

// WithEntities adds entities parameter
func (p *SendMessageParams) WithEntities(entities ...MessageEntity) *SendMessageParams {
	p.Entities = entities
	return p
}

// WithDisableWebPagePreview adds disable web page preview parameter
func (p *SendMessageParams) WithDisableWebPagePreview() *SendMessageParams {
	p.DisableWebPagePreview = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendMessageParams) WithDisableNotification() *SendMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendMessageParams) WithProtectContent() *SendMessageParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendMessageParams) WithReplyToMessageID(replyToMessageID int) *SendMessageParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendMessageParams) WithAllowSendingWithoutReply() *SendMessageParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendMessageParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendMessageParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *ForwardMessageParams) WithChatID(chatID ChatID) *ForwardMessageParams {
	p.ChatID = chatID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *ForwardMessageParams) WithFromChatID(fromChatID ChatID) *ForwardMessageParams {
	p.FromChatID = fromChatID
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *ForwardMessageParams) WithDisableNotification() *ForwardMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *ForwardMessageParams) WithProtectContent() *ForwardMessageParams {
	p.ProtectContent = true
	return p
}

// WithMessageID adds message ID parameter
func (p *ForwardMessageParams) WithMessageID(messageID int) *ForwardMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *CopyMessageParams) WithChatID(chatID ChatID) *CopyMessageParams {
	p.ChatID = chatID
	return p
}

// WithFromChatID adds from chat ID parameter
func (p *CopyMessageParams) WithFromChatID(fromChatID ChatID) *CopyMessageParams {
	p.FromChatID = fromChatID
	return p
}

// WithMessageID adds message ID parameter
func (p *CopyMessageParams) WithMessageID(messageID int) *CopyMessageParams {
	p.MessageID = messageID
	return p
}

// WithCaption adds caption parameter
func (p *CopyMessageParams) WithCaption(caption string) *CopyMessageParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *CopyMessageParams) WithParseMode(parseMode string) *CopyMessageParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *CopyMessageParams) WithCaptionEntities(captionEntities ...MessageEntity) *CopyMessageParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *CopyMessageParams) WithDisableNotification() *CopyMessageParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *CopyMessageParams) WithProtectContent() *CopyMessageParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *CopyMessageParams) WithReplyToMessageID(replyToMessageID int) *CopyMessageParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *CopyMessageParams) WithAllowSendingWithoutReply() *CopyMessageParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *CopyMessageParams) WithReplyMarkup(replyMarkup ReplyMarkup) *CopyMessageParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendPhotoParams) WithChatID(chatID ChatID) *SendPhotoParams {
	p.ChatID = chatID
	return p
}

// WithPhoto adds photo parameter
func (p *SendPhotoParams) WithPhoto(photo InputFile) *SendPhotoParams {
	p.Photo = photo
	return p
}

// WithCaption adds caption parameter
func (p *SendPhotoParams) WithCaption(caption string) *SendPhotoParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendPhotoParams) WithParseMode(parseMode string) *SendPhotoParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendPhotoParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendPhotoParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendPhotoParams) WithDisableNotification() *SendPhotoParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendPhotoParams) WithProtectContent() *SendPhotoParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendPhotoParams) WithReplyToMessageID(replyToMessageID int) *SendPhotoParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendPhotoParams) WithAllowSendingWithoutReply() *SendPhotoParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendPhotoParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendPhotoParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendAudioParams) WithChatID(chatID ChatID) *SendAudioParams {
	p.ChatID = chatID
	return p
}

// WithAudio adds audio parameter
func (p *SendAudioParams) WithAudio(audio InputFile) *SendAudioParams {
	p.Audio = audio
	return p
}

// WithCaption adds caption parameter
func (p *SendAudioParams) WithCaption(caption string) *SendAudioParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendAudioParams) WithParseMode(parseMode string) *SendAudioParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendAudioParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendAudioParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDuration adds duration parameter
func (p *SendAudioParams) WithDuration(duration int) *SendAudioParams {
	p.Duration = duration
	return p
}

// WithPerformer adds performer parameter
func (p *SendAudioParams) WithPerformer(performer string) *SendAudioParams {
	p.Performer = performer
	return p
}

// WithTitle adds title parameter
func (p *SendAudioParams) WithTitle(title string) *SendAudioParams {
	p.Title = title
	return p
}

// WithThumb adds thumb parameter
func (p *SendAudioParams) WithThumb(thumb *InputFile) *SendAudioParams {
	p.Thumb = thumb
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendAudioParams) WithDisableNotification() *SendAudioParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendAudioParams) WithProtectContent() *SendAudioParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendAudioParams) WithReplyToMessageID(replyToMessageID int) *SendAudioParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendAudioParams) WithAllowSendingWithoutReply() *SendAudioParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendAudioParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendAudioParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendDocumentParams) WithChatID(chatID ChatID) *SendDocumentParams {
	p.ChatID = chatID
	return p
}

// WithDocument adds document parameter
func (p *SendDocumentParams) WithDocument(document InputFile) *SendDocumentParams {
	p.Document = document
	return p
}

// WithThumb adds thumb parameter
func (p *SendDocumentParams) WithThumb(thumb *InputFile) *SendDocumentParams {
	p.Thumb = thumb
	return p
}

// WithCaption adds caption parameter
func (p *SendDocumentParams) WithCaption(caption string) *SendDocumentParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendDocumentParams) WithParseMode(parseMode string) *SendDocumentParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendDocumentParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendDocumentParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDisableContentTypeDetection adds disable content type detection parameter
func (p *SendDocumentParams) WithDisableContentTypeDetection() *SendDocumentParams {
	p.DisableContentTypeDetection = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendDocumentParams) WithDisableNotification() *SendDocumentParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendDocumentParams) WithProtectContent() *SendDocumentParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendDocumentParams) WithReplyToMessageID(replyToMessageID int) *SendDocumentParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendDocumentParams) WithAllowSendingWithoutReply() *SendDocumentParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendDocumentParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendDocumentParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVideoParams) WithChatID(chatID ChatID) *SendVideoParams {
	p.ChatID = chatID
	return p
}

// WithVideo adds video parameter
func (p *SendVideoParams) WithVideo(video InputFile) *SendVideoParams {
	p.Video = video
	return p
}

// WithDuration adds duration parameter
func (p *SendVideoParams) WithDuration(duration int) *SendVideoParams {
	p.Duration = duration
	return p
}

// WithWidth adds width parameter
func (p *SendVideoParams) WithWidth(width int) *SendVideoParams {
	p.Width = width
	return p
}

// WithHeight adds height parameter
func (p *SendVideoParams) WithHeight(height int) *SendVideoParams {
	p.Height = height
	return p
}

// WithThumb adds thumb parameter
func (p *SendVideoParams) WithThumb(thumb *InputFile) *SendVideoParams {
	p.Thumb = thumb
	return p
}

// WithCaption adds caption parameter
func (p *SendVideoParams) WithCaption(caption string) *SendVideoParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendVideoParams) WithParseMode(parseMode string) *SendVideoParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendVideoParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendVideoParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithSupportsStreaming adds supports streaming parameter
func (p *SendVideoParams) WithSupportsStreaming() *SendVideoParams {
	p.SupportsStreaming = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVideoParams) WithDisableNotification() *SendVideoParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVideoParams) WithProtectContent() *SendVideoParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendVideoParams) WithReplyToMessageID(replyToMessageID int) *SendVideoParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendVideoParams) WithAllowSendingWithoutReply() *SendVideoParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVideoParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVideoParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendAnimationParams) WithChatID(chatID ChatID) *SendAnimationParams {
	p.ChatID = chatID
	return p
}

// WithAnimation adds animation parameter
func (p *SendAnimationParams) WithAnimation(animation InputFile) *SendAnimationParams {
	p.Animation = animation
	return p
}

// WithDuration adds duration parameter
func (p *SendAnimationParams) WithDuration(duration int) *SendAnimationParams {
	p.Duration = duration
	return p
}

// WithWidth adds width parameter
func (p *SendAnimationParams) WithWidth(width int) *SendAnimationParams {
	p.Width = width
	return p
}

// WithHeight adds height parameter
func (p *SendAnimationParams) WithHeight(height int) *SendAnimationParams {
	p.Height = height
	return p
}

// WithThumb adds thumb parameter
func (p *SendAnimationParams) WithThumb(thumb *InputFile) *SendAnimationParams {
	p.Thumb = thumb
	return p
}

// WithCaption adds caption parameter
func (p *SendAnimationParams) WithCaption(caption string) *SendAnimationParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendAnimationParams) WithParseMode(parseMode string) *SendAnimationParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendAnimationParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendAnimationParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendAnimationParams) WithDisableNotification() *SendAnimationParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendAnimationParams) WithProtectContent() *SendAnimationParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendAnimationParams) WithReplyToMessageID(replyToMessageID int) *SendAnimationParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendAnimationParams) WithAllowSendingWithoutReply() *SendAnimationParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendAnimationParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendAnimationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVoiceParams) WithChatID(chatID ChatID) *SendVoiceParams {
	p.ChatID = chatID
	return p
}

// WithVoice adds voice parameter
func (p *SendVoiceParams) WithVoice(voice InputFile) *SendVoiceParams {
	p.Voice = voice
	return p
}

// WithCaption adds caption parameter
func (p *SendVoiceParams) WithCaption(caption string) *SendVoiceParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *SendVoiceParams) WithParseMode(parseMode string) *SendVoiceParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *SendVoiceParams) WithCaptionEntities(captionEntities ...MessageEntity) *SendVoiceParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithDuration adds duration parameter
func (p *SendVoiceParams) WithDuration(duration int) *SendVoiceParams {
	p.Duration = duration
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVoiceParams) WithDisableNotification() *SendVoiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVoiceParams) WithProtectContent() *SendVoiceParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendVoiceParams) WithReplyToMessageID(replyToMessageID int) *SendVoiceParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendVoiceParams) WithAllowSendingWithoutReply() *SendVoiceParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVoiceParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVoiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVideoNoteParams) WithChatID(chatID ChatID) *SendVideoNoteParams {
	p.ChatID = chatID
	return p
}

// WithVideoNote adds video note parameter
func (p *SendVideoNoteParams) WithVideoNote(videoNote InputFile) *SendVideoNoteParams {
	p.VideoNote = videoNote
	return p
}

// WithDuration adds duration parameter
func (p *SendVideoNoteParams) WithDuration(duration int) *SendVideoNoteParams {
	p.Duration = duration
	return p
}

// WithLength adds length parameter
func (p *SendVideoNoteParams) WithLength(length int) *SendVideoNoteParams {
	p.Length = length
	return p
}

// WithThumb adds thumb parameter
func (p *SendVideoNoteParams) WithThumb(thumb *InputFile) *SendVideoNoteParams {
	p.Thumb = thumb
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVideoNoteParams) WithDisableNotification() *SendVideoNoteParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVideoNoteParams) WithProtectContent() *SendVideoNoteParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendVideoNoteParams) WithReplyToMessageID(replyToMessageID int) *SendVideoNoteParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendVideoNoteParams) WithAllowSendingWithoutReply() *SendVideoNoteParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVideoNoteParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVideoNoteParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendMediaGroupParams) WithChatID(chatID ChatID) *SendMediaGroupParams {
	p.ChatID = chatID
	return p
}

// WithMedia adds media parameter
func (p *SendMediaGroupParams) WithMedia(media ...InputMedia) *SendMediaGroupParams {
	p.Media = media
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendMediaGroupParams) WithDisableNotification() *SendMediaGroupParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendMediaGroupParams) WithProtectContent() *SendMediaGroupParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendMediaGroupParams) WithReplyToMessageID(replyToMessageID int) *SendMediaGroupParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendMediaGroupParams) WithAllowSendingWithoutReply() *SendMediaGroupParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithChatID adds chat ID parameter
func (p *SendLocationParams) WithChatID(chatID ChatID) *SendLocationParams {
	p.ChatID = chatID
	return p
}

// WithLivePeriod adds live period parameter
func (p *SendLocationParams) WithLivePeriod(livePeriod int) *SendLocationParams {
	p.LivePeriod = livePeriod
	return p
}

// WithHeading adds heading parameter
func (p *SendLocationParams) WithHeading(heading int) *SendLocationParams {
	p.Heading = heading
	return p
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (p *SendLocationParams) WithProximityAlertRadius(proximityAlertRadius int) *SendLocationParams {
	p.ProximityAlertRadius = proximityAlertRadius
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendLocationParams) WithDisableNotification() *SendLocationParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendLocationParams) WithProtectContent() *SendLocationParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendLocationParams) WithReplyToMessageID(replyToMessageID int) *SendLocationParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendLocationParams) WithAllowSendingWithoutReply() *SendLocationParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendLocationParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageLiveLocationParams) WithChatID(chatID ChatID) *EditMessageLiveLocationParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageLiveLocationParams) WithMessageID(messageID int) *EditMessageLiveLocationParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageLiveLocationParams) WithInlineMessageID(inlineMessageID string) *EditMessageLiveLocationParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithHeading adds heading parameter
func (p *EditMessageLiveLocationParams) WithHeading(heading int) *EditMessageLiveLocationParams {
	p.Heading = heading
	return p
}

// WithProximityAlertRadius adds proximity alert radius parameter
func (p *EditMessageLiveLocationParams) WithProximityAlertRadius(proximityAlertRadius int,
) *EditMessageLiveLocationParams {
	p.ProximityAlertRadius = proximityAlertRadius
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageLiveLocationParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *EditMessageLiveLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *StopMessageLiveLocationParams) WithChatID(chatID ChatID) *StopMessageLiveLocationParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *StopMessageLiveLocationParams) WithMessageID(messageID int) *StopMessageLiveLocationParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *StopMessageLiveLocationParams) WithInlineMessageID(inlineMessageID string) *StopMessageLiveLocationParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *StopMessageLiveLocationParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *StopMessageLiveLocationParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendVenueParams) WithChatID(chatID ChatID) *SendVenueParams {
	p.ChatID = chatID
	return p
}

// WithTitle adds title parameter
func (p *SendVenueParams) WithTitle(title string) *SendVenueParams {
	p.Title = title
	return p
}

// WithAddress adds address parameter
func (p *SendVenueParams) WithAddress(address string) *SendVenueParams {
	p.Address = address
	return p
}

// WithFoursquareID adds foursquare ID parameter
func (p *SendVenueParams) WithFoursquareID(foursquareID string) *SendVenueParams {
	p.FoursquareID = foursquareID
	return p
}

// WithFoursquareType adds foursquare type parameter
func (p *SendVenueParams) WithFoursquareType(foursquareType string) *SendVenueParams {
	p.FoursquareType = foursquareType
	return p
}

// WithGooglePlaceID adds google place ID parameter
func (p *SendVenueParams) WithGooglePlaceID(googlePlaceID string) *SendVenueParams {
	p.GooglePlaceID = googlePlaceID
	return p
}

// WithGooglePlaceType adds google place type parameter
func (p *SendVenueParams) WithGooglePlaceType(googlePlaceType string) *SendVenueParams {
	p.GooglePlaceType = googlePlaceType
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendVenueParams) WithDisableNotification() *SendVenueParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendVenueParams) WithProtectContent() *SendVenueParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendVenueParams) WithReplyToMessageID(replyToMessageID int) *SendVenueParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendVenueParams) WithAllowSendingWithoutReply() *SendVenueParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendVenueParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendVenueParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendContactParams) WithChatID(chatID ChatID) *SendContactParams {
	p.ChatID = chatID
	return p
}

// WithPhoneNumber adds phone number parameter
func (p *SendContactParams) WithPhoneNumber(phoneNumber string) *SendContactParams {
	p.PhoneNumber = phoneNumber
	return p
}

// WithFirstName adds first name parameter
func (p *SendContactParams) WithFirstName(firstName string) *SendContactParams {
	p.FirstName = firstName
	return p
}

// WithLastName adds last name parameter
func (p *SendContactParams) WithLastName(lastName string) *SendContactParams {
	p.LastName = lastName
	return p
}

// WithVcard adds vcard parameter
func (p *SendContactParams) WithVcard(vcard string) *SendContactParams {
	p.Vcard = vcard
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendContactParams) WithDisableNotification() *SendContactParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendContactParams) WithProtectContent() *SendContactParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendContactParams) WithReplyToMessageID(replyToMessageID int) *SendContactParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendContactParams) WithAllowSendingWithoutReply() *SendContactParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendContactParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendContactParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendPollParams) WithChatID(chatID ChatID) *SendPollParams {
	p.ChatID = chatID
	return p
}

// WithQuestion adds question parameter
func (p *SendPollParams) WithQuestion(question string) *SendPollParams {
	p.Question = question
	return p
}

// WithOptions adds options parameter
func (p *SendPollParams) WithOptions(options ...string) *SendPollParams {
	p.Options = options
	return p
}

// WithIsAnonymous adds is anonymous parameter
func (p *SendPollParams) WithIsAnonymous() *SendPollParams {
	p.IsAnonymous = true
	return p
}

// WithType adds type parameter
func (p *SendPollParams) WithType(pollType string) *SendPollParams {
	p.Type = pollType
	return p
}

// WithAllowsMultipleAnswers adds allows multiple answers parameter
func (p *SendPollParams) WithAllowsMultipleAnswers() *SendPollParams {
	p.AllowsMultipleAnswers = true
	return p
}

// WithCorrectOptionID adds correct option ID parameter
func (p *SendPollParams) WithCorrectOptionID(correctOptionID int) *SendPollParams {
	p.CorrectOptionID = correctOptionID
	return p
}

// WithExplanation adds explanation parameter
func (p *SendPollParams) WithExplanation(explanation string) *SendPollParams {
	p.Explanation = explanation
	return p
}

// WithExplanationParseMode adds explanation parse mode parameter
func (p *SendPollParams) WithExplanationParseMode(explanationParseMode string) *SendPollParams {
	p.ExplanationParseMode = explanationParseMode
	return p
}

// WithExplanationEntities adds explanation entities parameter
func (p *SendPollParams) WithExplanationEntities(explanationEntities ...MessageEntity) *SendPollParams {
	p.ExplanationEntities = explanationEntities
	return p
}

// WithOpenPeriod adds open period parameter
func (p *SendPollParams) WithOpenPeriod(openPeriod int) *SendPollParams {
	p.OpenPeriod = openPeriod
	return p
}

// WithIsClosed adds is closed parameter
func (p *SendPollParams) WithIsClosed() *SendPollParams {
	p.IsClosed = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendPollParams) WithDisableNotification() *SendPollParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendPollParams) WithProtectContent() *SendPollParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendPollParams) WithReplyToMessageID(replyToMessageID int) *SendPollParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendPollParams) WithAllowSendingWithoutReply() *SendPollParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendPollParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendPollParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendDiceParams) WithChatID(chatID ChatID) *SendDiceParams {
	p.ChatID = chatID
	return p
}

// WithEmoji adds emoji parameter
func (p *SendDiceParams) WithEmoji(emoji string) *SendDiceParams {
	p.Emoji = emoji
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendDiceParams) WithDisableNotification() *SendDiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendDiceParams) WithProtectContent() *SendDiceParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendDiceParams) WithReplyToMessageID(replyToMessageID int) *SendDiceParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendDiceParams) WithAllowSendingWithoutReply() *SendDiceParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendDiceParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendDiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *SendChatActionParams) WithChatID(chatID ChatID) *SendChatActionParams {
	p.ChatID = chatID
	return p
}

// WithAction adds action parameter
func (p *SendChatActionParams) WithAction(action string) *SendChatActionParams {
	p.Action = action
	return p
}

// WithOffset adds offset parameter
func (p *GetUserProfilePhotosParams) WithOffset(offset int) *GetUserProfilePhotosParams {
	p.Offset = offset
	return p
}

// WithLimit adds limit parameter
func (p *GetUserProfilePhotosParams) WithLimit(limit int) *GetUserProfilePhotosParams {
	p.Limit = limit
	return p
}

// WithFileID adds file ID parameter
func (p *GetFileParams) WithFileID(fileID string) *GetFileParams {
	p.FileID = fileID
	return p
}

// WithChatID adds chat ID parameter
func (p *BanChatMemberParams) WithChatID(chatID ChatID) *BanChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithRevokeMessages adds revoke messages parameter
func (p *BanChatMemberParams) WithRevokeMessages() *BanChatMemberParams {
	p.RevokeMessages = true
	return p
}

// WithChatID adds chat ID parameter
func (p *UnbanChatMemberParams) WithChatID(chatID ChatID) *UnbanChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithOnlyIfBanned adds only if banned parameter
func (p *UnbanChatMemberParams) WithOnlyIfBanned() *UnbanChatMemberParams {
	p.OnlyIfBanned = true
	return p
}

// WithChatID adds chat ID parameter
func (p *RestrictChatMemberParams) WithChatID(chatID ChatID) *RestrictChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithPermissions adds permissions parameter
func (p *RestrictChatMemberParams) WithPermissions(permissions ChatPermissions) *RestrictChatMemberParams {
	p.Permissions = permissions
	return p
}

// WithChatID adds chat ID parameter
func (p *PromoteChatMemberParams) WithChatID(chatID ChatID) *PromoteChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithIsAnonymous adds is anonymous parameter
func (p *PromoteChatMemberParams) WithIsAnonymous() *PromoteChatMemberParams {
	p.IsAnonymous = true
	return p
}

// WithCanManageChat adds can manage chat parameter
func (p *PromoteChatMemberParams) WithCanManageChat() *PromoteChatMemberParams {
	p.CanManageChat = true
	return p
}

// WithCanPostMessages adds can post messages parameter
func (p *PromoteChatMemberParams) WithCanPostMessages() *PromoteChatMemberParams {
	p.CanPostMessages = true
	return p
}

// WithCanEditMessages adds can edit messages parameter
func (p *PromoteChatMemberParams) WithCanEditMessages() *PromoteChatMemberParams {
	p.CanEditMessages = true
	return p
}

// WithCanDeleteMessages adds can delete messages parameter
func (p *PromoteChatMemberParams) WithCanDeleteMessages() *PromoteChatMemberParams {
	p.CanDeleteMessages = true
	return p
}

// WithCanManageVoiceChats adds can manage voice chats parameter
func (p *PromoteChatMemberParams) WithCanManageVoiceChats() *PromoteChatMemberParams {
	p.CanManageVoiceChats = true
	return p
}

// WithCanRestrictMembers adds can restrict members parameter
func (p *PromoteChatMemberParams) WithCanRestrictMembers() *PromoteChatMemberParams {
	p.CanRestrictMembers = true
	return p
}

// WithCanPromoteMembers adds can promote members parameter
func (p *PromoteChatMemberParams) WithCanPromoteMembers() *PromoteChatMemberParams {
	p.CanPromoteMembers = true
	return p
}

// WithCanChangeInfo adds can change info parameter
func (p *PromoteChatMemberParams) WithCanChangeInfo() *PromoteChatMemberParams {
	p.CanChangeInfo = true
	return p
}

// WithCanInviteUsers adds can invite users parameter
func (p *PromoteChatMemberParams) WithCanInviteUsers() *PromoteChatMemberParams {
	p.CanInviteUsers = true
	return p
}

// WithCanPinMessages adds can pin messages parameter
func (p *PromoteChatMemberParams) WithCanPinMessages() *PromoteChatMemberParams {
	p.CanPinMessages = true
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatAdministratorCustomTitleParams) WithChatID(chatID ChatID) *SetChatAdministratorCustomTitleParams {
	p.ChatID = chatID
	return p
}

// WithCustomTitle adds custom title parameter
func (p *SetChatAdministratorCustomTitleParams) WithCustomTitle(customTitle string,
) *SetChatAdministratorCustomTitleParams {
	p.CustomTitle = customTitle
	return p
}

// WithChatID adds chat ID parameter
func (p *BanChatSenderChatParams) WithChatID(chatID ChatID) *BanChatSenderChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnbanChatSenderChatParams) WithChatID(chatID ChatID) *UnbanChatSenderChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatPermissionsParams) WithChatID(chatID ChatID) *SetChatPermissionsParams {
	p.ChatID = chatID
	return p
}

// WithPermissions adds permissions parameter
func (p *SetChatPermissionsParams) WithPermissions(permissions ChatPermissions) *SetChatPermissionsParams {
	p.Permissions = permissions
	return p
}

// WithChatID adds chat ID parameter
func (p *ExportChatInviteLinkParams) WithChatID(chatID ChatID) *ExportChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *CreateChatInviteLinkParams) WithChatID(chatID ChatID) *CreateChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithName adds name parameter
func (p *CreateChatInviteLinkParams) WithName(name string) *CreateChatInviteLinkParams {
	p.Name = name
	return p
}

// WithMemberLimit adds member limit parameter
func (p *CreateChatInviteLinkParams) WithMemberLimit(memberLimit int) *CreateChatInviteLinkParams {
	p.MemberLimit = memberLimit
	return p
}

// WithCreatesJoinRequest adds creates join request parameter
func (p *CreateChatInviteLinkParams) WithCreatesJoinRequest() *CreateChatInviteLinkParams {
	p.CreatesJoinRequest = true
	return p
}

// WithChatID adds chat ID parameter
func (p *EditChatInviteLinkParams) WithChatID(chatID ChatID) *EditChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithInviteLink adds invite link parameter
func (p *EditChatInviteLinkParams) WithInviteLink(inviteLink string) *EditChatInviteLinkParams {
	p.InviteLink = inviteLink
	return p
}

// WithName adds name parameter
func (p *EditChatInviteLinkParams) WithName(name string) *EditChatInviteLinkParams {
	p.Name = name
	return p
}

// WithMemberLimit adds member limit parameter
func (p *EditChatInviteLinkParams) WithMemberLimit(memberLimit int) *EditChatInviteLinkParams {
	p.MemberLimit = memberLimit
	return p
}

// WithCreatesJoinRequest adds creates join request parameter
func (p *EditChatInviteLinkParams) WithCreatesJoinRequest() *EditChatInviteLinkParams {
	p.CreatesJoinRequest = true
	return p
}

// WithChatID adds chat ID parameter
func (p *RevokeChatInviteLinkParams) WithChatID(chatID ChatID) *RevokeChatInviteLinkParams {
	p.ChatID = chatID
	return p
}

// WithInviteLink adds invite link parameter
func (p *RevokeChatInviteLinkParams) WithInviteLink(inviteLink string) *RevokeChatInviteLinkParams {
	p.InviteLink = inviteLink
	return p
}

// WithChatID adds chat ID parameter
func (p *ApproveChatJoinRequestParams) WithChatID(chatID ChatID) *ApproveChatJoinRequestParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *DeclineChatJoinRequestParams) WithChatID(chatID ChatID) *DeclineChatJoinRequestParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatPhotoParams) WithChatID(chatID ChatID) *SetChatPhotoParams {
	p.ChatID = chatID
	return p
}

// WithPhoto adds photo parameter
func (p *SetChatPhotoParams) WithPhoto(photo InputFile) *SetChatPhotoParams {
	p.Photo = photo
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteChatPhotoParams) WithChatID(chatID ChatID) *DeleteChatPhotoParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatTitleParams) WithChatID(chatID ChatID) *SetChatTitleParams {
	p.ChatID = chatID
	return p
}

// WithTitle adds title parameter
func (p *SetChatTitleParams) WithTitle(title string) *SetChatTitleParams {
	p.Title = title
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatDescriptionParams) WithChatID(chatID ChatID) *SetChatDescriptionParams {
	p.ChatID = chatID
	return p
}

// WithDescription adds description parameter
func (p *SetChatDescriptionParams) WithDescription(description string) *SetChatDescriptionParams {
	p.Description = description
	return p
}

// WithChatID adds chat ID parameter
func (p *PinChatMessageParams) WithChatID(chatID ChatID) *PinChatMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *PinChatMessageParams) WithMessageID(messageID int) *PinChatMessageParams {
	p.MessageID = messageID
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *PinChatMessageParams) WithDisableNotification() *PinChatMessageParams {
	p.DisableNotification = true
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinChatMessageParams) WithChatID(chatID ChatID) *UnpinChatMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *UnpinChatMessageParams) WithMessageID(messageID int) *UnpinChatMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *UnpinAllChatMessagesParams) WithChatID(chatID ChatID) *UnpinAllChatMessagesParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *LeaveChatParams) WithChatID(chatID ChatID) *LeaveChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatParams) WithChatID(chatID ChatID) *GetChatParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatAdministratorsParams) WithChatID(chatID ChatID) *GetChatAdministratorsParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatMemberCountParams) WithChatID(chatID ChatID) *GetChatMemberCountParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *GetChatMemberParams) WithChatID(chatID ChatID) *GetChatMemberParams {
	p.ChatID = chatID
	return p
}

// WithChatID adds chat ID parameter
func (p *SetChatStickerSetParams) WithChatID(chatID ChatID) *SetChatStickerSetParams {
	p.ChatID = chatID
	return p
}

// WithStickerSetName adds sticker set name parameter
func (p *SetChatStickerSetParams) WithStickerSetName(stickerSetName string) *SetChatStickerSetParams {
	p.StickerSetName = stickerSetName
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteChatStickerSetParams) WithChatID(chatID ChatID) *DeleteChatStickerSetParams {
	p.ChatID = chatID
	return p
}

// WithCallbackQueryID adds callback query ID parameter
func (p *AnswerCallbackQueryParams) WithCallbackQueryID(callbackQueryID string) *AnswerCallbackQueryParams {
	p.CallbackQueryID = callbackQueryID
	return p
}

// WithText adds text parameter
func (p *AnswerCallbackQueryParams) WithText(text string) *AnswerCallbackQueryParams {
	p.Text = text
	return p
}

// WithShowAlert adds show alert parameter
func (p *AnswerCallbackQueryParams) WithShowAlert() *AnswerCallbackQueryParams {
	p.ShowAlert = true
	return p
}

// WithURL adds URL parameter
func (p *AnswerCallbackQueryParams) WithURL(url string) *AnswerCallbackQueryParams {
	p.URL = url
	return p
}

// WithCacheTime adds cache time parameter
func (p *AnswerCallbackQueryParams) WithCacheTime(cacheTime int) *AnswerCallbackQueryParams {
	p.CacheTime = cacheTime
	return p
}

// WithCommands adds commands parameter
func (p *SetMyCommandsParams) WithCommands(commands ...BotCommand) *SetMyCommandsParams {
	p.Commands = commands
	return p
}

// WithScope adds scope parameter
func (p *SetMyCommandsParams) WithScope(scope BotCommandScope) *SetMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *SetMyCommandsParams) WithLanguageCode(languageCode string) *SetMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithScope adds scope parameter
func (p *DeleteMyCommandsParams) WithScope(scope BotCommandScope) *DeleteMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *DeleteMyCommandsParams) WithLanguageCode(languageCode string) *DeleteMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithScope adds scope parameter
func (p *GetMyCommandsParams) WithScope(scope BotCommandScope) *GetMyCommandsParams {
	p.Scope = scope
	return p
}

// WithLanguageCode adds language code parameter
func (p *GetMyCommandsParams) WithLanguageCode(languageCode string) *GetMyCommandsParams {
	p.LanguageCode = languageCode
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageTextParams) WithChatID(chatID ChatID) *EditMessageTextParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageTextParams) WithMessageID(messageID int) *EditMessageTextParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageTextParams) WithInlineMessageID(inlineMessageID string) *EditMessageTextParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithText adds text parameter
func (p *EditMessageTextParams) WithText(text string) *EditMessageTextParams {
	p.Text = text
	return p
}

// WithParseMode adds parse mode parameter
func (p *EditMessageTextParams) WithParseMode(parseMode string) *EditMessageTextParams {
	p.ParseMode = parseMode
	return p
}

// WithEntities adds entities parameter
func (p *EditMessageTextParams) WithEntities(entities ...MessageEntity) *EditMessageTextParams {
	p.Entities = entities
	return p
}

// WithDisableWebPagePreview adds disable web page preview parameter
func (p *EditMessageTextParams) WithDisableWebPagePreview() *EditMessageTextParams {
	p.DisableWebPagePreview = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageTextParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageTextParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageCaptionParams) WithChatID(chatID ChatID) *EditMessageCaptionParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageCaptionParams) WithMessageID(messageID int) *EditMessageCaptionParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageCaptionParams) WithInlineMessageID(inlineMessageID string) *EditMessageCaptionParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithCaption adds caption parameter
func (p *EditMessageCaptionParams) WithCaption(caption string) *EditMessageCaptionParams {
	p.Caption = caption
	return p
}

// WithParseMode adds parse mode parameter
func (p *EditMessageCaptionParams) WithParseMode(parseMode string) *EditMessageCaptionParams {
	p.ParseMode = parseMode
	return p
}

// WithCaptionEntities adds caption entities parameter
func (p *EditMessageCaptionParams) WithCaptionEntities(captionEntities ...MessageEntity) *EditMessageCaptionParams {
	p.CaptionEntities = captionEntities
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageCaptionParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageCaptionParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageMediaParams) WithChatID(chatID ChatID) *EditMessageMediaParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageMediaParams) WithMessageID(messageID int) *EditMessageMediaParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageMediaParams) WithInlineMessageID(inlineMessageID string) *EditMessageMediaParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithMedia adds media parameter
func (p *EditMessageMediaParams) WithMedia(media InputMedia) *EditMessageMediaParams {
	p.Media = media
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageMediaParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *EditMessageMediaParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *EditMessageReplyMarkupParams) WithChatID(chatID ChatID) *EditMessageReplyMarkupParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *EditMessageReplyMarkupParams) WithMessageID(messageID int) *EditMessageReplyMarkupParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *EditMessageReplyMarkupParams) WithInlineMessageID(inlineMessageID string) *EditMessageReplyMarkupParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *EditMessageReplyMarkupParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup,
) *EditMessageReplyMarkupParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *StopPollParams) WithChatID(chatID ChatID) *StopPollParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *StopPollParams) WithMessageID(messageID int) *StopPollParams {
	p.MessageID = messageID
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *StopPollParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *StopPollParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithChatID adds chat ID parameter
func (p *DeleteMessageParams) WithChatID(chatID ChatID) *DeleteMessageParams {
	p.ChatID = chatID
	return p
}

// WithMessageID adds message ID parameter
func (p *DeleteMessageParams) WithMessageID(messageID int) *DeleteMessageParams {
	p.MessageID = messageID
	return p
}

// WithChatID adds chat ID parameter
func (p *SendStickerParams) WithChatID(chatID ChatID) *SendStickerParams {
	p.ChatID = chatID
	return p
}

// WithSticker adds sticker parameter
func (p *SendStickerParams) WithSticker(sticker InputFile) *SendStickerParams {
	p.Sticker = sticker
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendStickerParams) WithDisableNotification() *SendStickerParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendStickerParams) WithProtectContent() *SendStickerParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendStickerParams) WithReplyToMessageID(replyToMessageID int) *SendStickerParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendStickerParams) WithAllowSendingWithoutReply() *SendStickerParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendStickerParams) WithReplyMarkup(replyMarkup ReplyMarkup) *SendStickerParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithName adds name parameter
func (p *GetStickerSetParams) WithName(name string) *GetStickerSetParams {
	p.Name = name
	return p
}

// WithPngSticker adds png sticker parameter
func (p *UploadStickerFileParams) WithPngSticker(pngSticker InputFile) *UploadStickerFileParams {
	p.PngSticker = pngSticker
	return p
}

// WithName adds name parameter
func (p *CreateNewStickerSetParams) WithName(name string) *CreateNewStickerSetParams {
	p.Name = name
	return p
}

// WithTitle adds title parameter
func (p *CreateNewStickerSetParams) WithTitle(title string) *CreateNewStickerSetParams {
	p.Title = title
	return p
}

// WithPngSticker adds png sticker parameter
func (p *CreateNewStickerSetParams) WithPngSticker(pngSticker *InputFile) *CreateNewStickerSetParams {
	p.PngSticker = pngSticker
	return p
}

// WithTgsSticker adds tgs sticker parameter
func (p *CreateNewStickerSetParams) WithTgsSticker(tgsSticker *InputFile) *CreateNewStickerSetParams {
	p.TgsSticker = tgsSticker
	return p
}

// WithWebmSticker adds webm sticker parameter
func (p *CreateNewStickerSetParams) WithWebmSticker(webmSticker *InputFile) *CreateNewStickerSetParams {
	p.WebmSticker = webmSticker
	return p
}

// WithEmojis adds emojis parameter
func (p *CreateNewStickerSetParams) WithEmojis(emojis string) *CreateNewStickerSetParams {
	p.Emojis = emojis
	return p
}

// WithContainsMasks adds contains masks parameter
func (p *CreateNewStickerSetParams) WithContainsMasks() *CreateNewStickerSetParams {
	p.ContainsMasks = true
	return p
}

// WithMaskPosition adds mask position parameter
func (p *CreateNewStickerSetParams) WithMaskPosition(maskPosition *MaskPosition) *CreateNewStickerSetParams {
	p.MaskPosition = maskPosition
	return p
}

// WithName adds name parameter
func (p *AddStickerToSetParams) WithName(name string) *AddStickerToSetParams {
	p.Name = name
	return p
}

// WithPngSticker adds png sticker parameter
func (p *AddStickerToSetParams) WithPngSticker(pngSticker *InputFile) *AddStickerToSetParams {
	p.PngSticker = pngSticker
	return p
}

// WithTgsSticker adds tgs sticker parameter
func (p *AddStickerToSetParams) WithTgsSticker(tgsSticker *InputFile) *AddStickerToSetParams {
	p.TgsSticker = tgsSticker
	return p
}

// WithWebmSticker adds webm sticker parameter
func (p *AddStickerToSetParams) WithWebmSticker(webmSticker *InputFile) *AddStickerToSetParams {
	p.WebmSticker = webmSticker
	return p
}

// WithEmojis adds emojis parameter
func (p *AddStickerToSetParams) WithEmojis(emojis string) *AddStickerToSetParams {
	p.Emojis = emojis
	return p
}

// WithMaskPosition adds mask position parameter
func (p *AddStickerToSetParams) WithMaskPosition(maskPosition *MaskPosition) *AddStickerToSetParams {
	p.MaskPosition = maskPosition
	return p
}

// WithSticker adds sticker parameter
func (p *SetStickerPositionInSetParams) WithSticker(sticker string) *SetStickerPositionInSetParams {
	p.Sticker = sticker
	return p
}

// WithPosition adds position parameter
func (p *SetStickerPositionInSetParams) WithPosition(position int) *SetStickerPositionInSetParams {
	p.Position = position
	return p
}

// WithSticker adds sticker parameter
func (p *DeleteStickerFromSetParams) WithSticker(sticker string) *DeleteStickerFromSetParams {
	p.Sticker = sticker
	return p
}

// WithName adds name parameter
func (p *SetStickerSetThumbParams) WithName(name string) *SetStickerSetThumbParams {
	p.Name = name
	return p
}

// WithThumb adds thumb parameter
func (p *SetStickerSetThumbParams) WithThumb(thumb *InputFile) *SetStickerSetThumbParams {
	p.Thumb = thumb
	return p
}

// WithInlineQueryID adds inline query ID parameter
func (p *AnswerInlineQueryParams) WithInlineQueryID(inlineQueryID string) *AnswerInlineQueryParams {
	p.InlineQueryID = inlineQueryID
	return p
}

// WithResults adds results parameter
func (p *AnswerInlineQueryParams) WithResults(results ...InlineQueryResult) *AnswerInlineQueryParams {
	p.Results = results
	return p
}

// WithCacheTime adds cache time parameter
func (p *AnswerInlineQueryParams) WithCacheTime(cacheTime int) *AnswerInlineQueryParams {
	p.CacheTime = cacheTime
	return p
}

// WithIsPersonal adds is personal parameter
func (p *AnswerInlineQueryParams) WithIsPersonal() *AnswerInlineQueryParams {
	p.IsPersonal = true
	return p
}

// WithNextOffset adds next offset parameter
func (p *AnswerInlineQueryParams) WithNextOffset(nextOffset string) *AnswerInlineQueryParams {
	p.NextOffset = nextOffset
	return p
}

// WithSwitchPmText adds switch pm text parameter
func (p *AnswerInlineQueryParams) WithSwitchPmText(switchPmText string) *AnswerInlineQueryParams {
	p.SwitchPmText = switchPmText
	return p
}

// WithSwitchPmParameter adds switch pm parameter parameter
func (p *AnswerInlineQueryParams) WithSwitchPmParameter(switchPmParameter string) *AnswerInlineQueryParams {
	p.SwitchPmParameter = switchPmParameter
	return p
}

// WithChatID adds chat ID parameter
func (p *SendInvoiceParams) WithChatID(chatID ChatID) *SendInvoiceParams {
	p.ChatID = chatID
	return p
}

// WithTitle adds title parameter
func (p *SendInvoiceParams) WithTitle(title string) *SendInvoiceParams {
	p.Title = title
	return p
}

// WithDescription adds description parameter
func (p *SendInvoiceParams) WithDescription(description string) *SendInvoiceParams {
	p.Description = description
	return p
}

// WithPayload adds payload parameter
func (p *SendInvoiceParams) WithPayload(payload string) *SendInvoiceParams {
	p.Payload = payload
	return p
}

// WithProviderToken adds provider token parameter
func (p *SendInvoiceParams) WithProviderToken(providerToken string) *SendInvoiceParams {
	p.ProviderToken = providerToken
	return p
}

// WithCurrency adds currency parameter
func (p *SendInvoiceParams) WithCurrency(currency string) *SendInvoiceParams {
	p.Currency = currency
	return p
}

// WithPrices adds prices parameter
func (p *SendInvoiceParams) WithPrices(prices ...LabeledPrice) *SendInvoiceParams {
	p.Prices = prices
	return p
}

// WithMaxTipAmount adds max tip amount parameter
func (p *SendInvoiceParams) WithMaxTipAmount(maxTipAmount int) *SendInvoiceParams {
	p.MaxTipAmount = maxTipAmount
	return p
}

// WithSuggestedTipAmounts adds suggested tip amounts parameter
func (p *SendInvoiceParams) WithSuggestedTipAmounts(suggestedTipAmounts ...int) *SendInvoiceParams {
	p.SuggestedTipAmounts = suggestedTipAmounts
	return p
}

// WithStartParameter adds start parameter parameter
func (p *SendInvoiceParams) WithStartParameter(startParameter string) *SendInvoiceParams {
	p.StartParameter = startParameter
	return p
}

// WithProviderData adds provider data parameter
func (p *SendInvoiceParams) WithProviderData(providerData string) *SendInvoiceParams {
	p.ProviderData = providerData
	return p
}

// WithPhotoURL adds photo URL parameter
func (p *SendInvoiceParams) WithPhotoURL(photoURL string) *SendInvoiceParams {
	p.PhotoURL = photoURL
	return p
}

// WithPhotoSize adds photo size parameter
func (p *SendInvoiceParams) WithPhotoSize(photoSize int) *SendInvoiceParams {
	p.PhotoSize = photoSize
	return p
}

// WithPhotoWidth adds photo width parameter
func (p *SendInvoiceParams) WithPhotoWidth(photoWidth int) *SendInvoiceParams {
	p.PhotoWidth = photoWidth
	return p
}

// WithPhotoHeight adds photo height parameter
func (p *SendInvoiceParams) WithPhotoHeight(photoHeight int) *SendInvoiceParams {
	p.PhotoHeight = photoHeight
	return p
}

// WithNeedName adds need name parameter
func (p *SendInvoiceParams) WithNeedName() *SendInvoiceParams {
	p.NeedName = true
	return p
}

// WithNeedPhoneNumber adds need phone number parameter
func (p *SendInvoiceParams) WithNeedPhoneNumber() *SendInvoiceParams {
	p.NeedPhoneNumber = true
	return p
}

// WithNeedEmail adds need email parameter
func (p *SendInvoiceParams) WithNeedEmail() *SendInvoiceParams {
	p.NeedEmail = true
	return p
}

// WithNeedShippingAddress adds need shipping address parameter
func (p *SendInvoiceParams) WithNeedShippingAddress() *SendInvoiceParams {
	p.NeedShippingAddress = true
	return p
}

// WithSendPhoneNumberToProvider adds send phone number to provider parameter
func (p *SendInvoiceParams) WithSendPhoneNumberToProvider() *SendInvoiceParams {
	p.SendPhoneNumberToProvider = true
	return p
}

// WithSendEmailToProvider adds send email to provider parameter
func (p *SendInvoiceParams) WithSendEmailToProvider() *SendInvoiceParams {
	p.SendEmailToProvider = true
	return p
}

// WithIsFlexible adds is flexible parameter
func (p *SendInvoiceParams) WithIsFlexible() *SendInvoiceParams {
	p.IsFlexible = true
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendInvoiceParams) WithDisableNotification() *SendInvoiceParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendInvoiceParams) WithProtectContent() *SendInvoiceParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendInvoiceParams) WithReplyToMessageID(replyToMessageID int) *SendInvoiceParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendInvoiceParams) WithAllowSendingWithoutReply() *SendInvoiceParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendInvoiceParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *SendInvoiceParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithShippingQueryID adds shipping query ID parameter
func (p *AnswerShippingQueryParams) WithShippingQueryID(shippingQueryID string) *AnswerShippingQueryParams {
	p.ShippingQueryID = shippingQueryID
	return p
}

// WithOk adds ok parameter
func (p *AnswerShippingQueryParams) WithOk() *AnswerShippingQueryParams {
	p.Ok = true
	return p
}

// WithShippingOptions adds shipping options parameter
func (p *AnswerShippingQueryParams) WithShippingOptions(shippingOptions ...ShippingOption) *AnswerShippingQueryParams {
	p.ShippingOptions = shippingOptions
	return p
}

// WithErrorMessage adds error message parameter
func (p *AnswerShippingQueryParams) WithErrorMessage(errorMessage string) *AnswerShippingQueryParams {
	p.ErrorMessage = errorMessage
	return p
}

// WithPreCheckoutQueryID adds pre checkout query ID parameter
func (p *AnswerPreCheckoutQueryParams) WithPreCheckoutQueryID(preCheckoutQueryID string) *AnswerPreCheckoutQueryParams {
	p.PreCheckoutQueryID = preCheckoutQueryID
	return p
}

// WithOk adds ok parameter
func (p *AnswerPreCheckoutQueryParams) WithOk() *AnswerPreCheckoutQueryParams {
	p.Ok = true
	return p
}

// WithErrorMessage adds error message parameter
func (p *AnswerPreCheckoutQueryParams) WithErrorMessage(errorMessage string) *AnswerPreCheckoutQueryParams {
	p.ErrorMessage = errorMessage
	return p
}

// WithErrors adds errors parameter
func (p *SetPassportDataErrorsParams) WithErrors(errors ...PassportElementError) *SetPassportDataErrorsParams {
	p.Errors = errors
	return p
}

// WithGameShortName adds game short name parameter
func (p *SendGameParams) WithGameShortName(gameShortName string) *SendGameParams {
	p.GameShortName = gameShortName
	return p
}

// WithDisableNotification adds disable notification parameter
func (p *SendGameParams) WithDisableNotification() *SendGameParams {
	p.DisableNotification = true
	return p
}

// WithProtectContent adds protect content parameter
func (p *SendGameParams) WithProtectContent() *SendGameParams {
	p.ProtectContent = true
	return p
}

// WithReplyToMessageID adds reply to message ID parameter
func (p *SendGameParams) WithReplyToMessageID(replyToMessageID int) *SendGameParams {
	p.ReplyToMessageID = replyToMessageID
	return p
}

// WithAllowSendingWithoutReply adds allow sending without reply parameter
func (p *SendGameParams) WithAllowSendingWithoutReply() *SendGameParams {
	p.AllowSendingWithoutReply = true
	return p
}

// WithReplyMarkup adds reply markup parameter
func (p *SendGameParams) WithReplyMarkup(replyMarkup *InlineKeyboardMarkup) *SendGameParams {
	p.ReplyMarkup = replyMarkup
	return p
}

// WithScore adds score parameter
func (p *SetGameScoreParams) WithScore(score int) *SetGameScoreParams {
	p.Score = score
	return p
}

// WithForce adds force parameter
func (p *SetGameScoreParams) WithForce() *SetGameScoreParams {
	p.Force = true
	return p
}

// WithDisableEditMessage adds disable edit message parameter
func (p *SetGameScoreParams) WithDisableEditMessage() *SetGameScoreParams {
	p.DisableEditMessage = true
	return p
}

// WithMessageID adds message ID parameter
func (p *SetGameScoreParams) WithMessageID(messageID int) *SetGameScoreParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *SetGameScoreParams) WithInlineMessageID(inlineMessageID string) *SetGameScoreParams {
	p.InlineMessageID = inlineMessageID
	return p
}

// WithMessageID adds message ID parameter
func (p *GetGameHighScoresParams) WithMessageID(messageID int) *GetGameHighScoresParams {
	p.MessageID = messageID
	return p
}

// WithInlineMessageID adds inline message ID parameter
func (p *GetGameHighScoresParams) WithInlineMessageID(inlineMessageID string) *GetGameHighScoresParams {
	p.InlineMessageID = inlineMessageID
	return p
}
