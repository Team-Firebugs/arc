/*
 * Arc - Copyleft of Simone 'evilsocket' Margaritelli.
 * evilsocket at protonmail dot com
 * https://www.evilsocket.net/
 *
 * See LICENSE.
 */
package events

import (
	"fmt"
	"github.com/evilsocket/arc/arcd/db"
	"html"
)

func Login(successful bool, address string, username string, password string) Event {
	if successful {
		desc := fmt.Sprintf("Address %s successfully logged into the Arc server.", address)
		return New("login_ok", "Successful login.", desc)
	} else {
		desc := fmt.Sprintf("Address <b>%s</b> tried to log into the Arc server with username <b>%s</b> and password <b>%s</b>.",
			html.EscapeString(address),
			html.EscapeString(username),
			html.EscapeString(password))
		return New("login_ko", "Failed login attempt.", desc)
	}
}

func InvalidToken(address, auth string, err error) Event {
	title := "Invalid token authentication."
	reason := ""
	if err != nil {
		reason = fmt.Sprintf(" (reason: %s)", err.Error())
	}

	desc := fmt.Sprintf("Address <b>%s</b> tried to authenticate with an invalid token '%s'%s.",
		html.EscapeString(address),
		html.EscapeString(auth),
		reason)

	return New("token_ko", title, desc)
}

func RecordExpired(r *db.Record) Event {
	meta := r.Meta()

	title := fmt.Sprintf("'%s' just expired.", meta.Title)
	compressed := ""
	deleted := ""

	if meta.Compressed {
		compressed = "(and gzipped) "
	}

	if meta.Prune {
		deleted = "It has been deleted from the system."
	}

	desc := fmt.Sprintf("The record <b>%s</b> which was created on %s and updated on %s just expired, it was made of %d bytes of %s encrypted %sdata.%s",
		html.EscapeString(meta.Title),
		meta.CreatedAt.Format("Mon Jan 2 15:04:05 2006"),
		meta.UpdatedAt.Format("Mon Jan 2 15:04:05 2006"),
		meta.Size,
		meta.Encryption,
		compressed,
		deleted)

	return New("record_expired", title, desc)
}
