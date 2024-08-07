/**
 * Webhook Type Definition
 * Webhook event definition of the LINE Messaging API
 *
 * The version of the OpenAPI document: 1.0.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package webhook

import (
	"encoding/json"
	"fmt"
)

// FollowEvent
// Event object for when your LINE Official Account is added as a friend (or unblocked). You can reply to follow events.

type FollowEvent struct {
	Event

	/**
	 * Get Source
	 */
	Source SourceInterface `json:"source,omitempty"`

	/**
	 * Time of the event in milliseconds. (Required)
	 */
	Timestamp int64 `json:"timestamp"`

	/**
	 * Get Mode
	 */
	Mode EventMode `json:"mode"`

	/**
	 * Webhook Event ID. An ID that uniquely identifies a webhook event. This is a string in ULID format. (Required)
	 */
	WebhookEventId string `json:"webhookEventId"`

	/**
	 * Get DeliveryContext
	 */
	DeliveryContext *DeliveryContext `json:"deliveryContext"`

	/**
	 * Reply token used to send reply message to this event (Required)
	 */
	ReplyToken string `json:"replyToken"`
}

func (cr *FollowEvent) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["type"] != nil {

		err = json.Unmarshal(raw["type"], &cr.Type)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(Type): %w", err)
		}

	}

	if raw["source"] != nil {

		if rawsource, ok := raw["source"]; ok && rawsource != nil {
			Source, err := UnmarshalSource(rawsource)
			if err != nil {
				return fmt.Errorf("JSON parse error in Source(discriminator): %w", err)
			}
			cr.Source = Source
		}

	}

	if raw["timestamp"] != nil {

		err = json.Unmarshal(raw["timestamp"], &cr.Timestamp)
		if err != nil {
			return fmt.Errorf("JSON parse error in int64(Timestamp): %w", err)
		}

	}

	if raw["mode"] != nil {

		err = json.Unmarshal(raw["mode"], &cr.Mode)
		if err != nil {
			return fmt.Errorf("JSON parse error in EventMode(Mode): %w", err)
		}

	}

	if raw["webhookEventId"] != nil {

		err = json.Unmarshal(raw["webhookEventId"], &cr.WebhookEventId)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(WebhookEventId): %w", err)
		}

	}

	if raw["deliveryContext"] != nil {

		err = json.Unmarshal(raw["deliveryContext"], &cr.DeliveryContext)
		if err != nil {
			return fmt.Errorf("JSON parse error in DeliveryContext(DeliveryContext): %w", err)
		}

	}

	if raw["replyToken"] != nil {

		err = json.Unmarshal(raw["replyToken"], &cr.ReplyToken)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(ReplyToken): %w", err)
		}

	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the FollowEvent struct.
func (r *FollowEvent) MarshalJSON() ([]byte, error) {

	r.Source = setDiscriminatorPropertySource(r.Source)

	type Alias FollowEvent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "follow",
	})
}
