package main

import (
	"os"
	"strings"
	"time"
)

const GUITAR_STRINGS = "EADGBE"
const MUSIC_KEYS = "CGDAE"  // focus on common guitar keys

func MusicKeysSlice() []string {
	return strings.Split(MUSIC_KEYS, "")
}


func EmitKey(doc DocWriter, key string, isMajor bool) {
	keyName := key
	if isMajor {
		keyName += " major"
	} else {
		keyName += " minor"
	}

	doc.EmitHeader("TODO Key %s", keyName)

	if isMajor {
		doc.EmitListCheckboxLn("Find I, iii and V for %s", keyName)
	} else {
		doc.EmitListCheckboxLn("Find i, III and v for %s", keyName)
	}

	doc.Depth++

	// create a daily practice of three adjacent string triads,
	// starting with the highest sounding three strings on the guitar,
	// and moving down
	for gStringLowest := 3; gStringLowest != -1; gStringLowest-- {
		gStringHighest := gStringLowest + 3

		gStringSet := GUITAR_STRINGS[gStringLowest:gStringHighest]
		doc.EmitHeader("TODO %s - Strings '%s'\n", keyName, gStringSet)

		doc.EmitListCheckboxLn("Play each triad up the board on strings %d-%d",
			gStringLowest+1, gStringHighest)
		doc.EmitListCheckboxLn("Observe the inversion and root for each triad")
		doc.EmitListCheckboxLn("Compose a two-chord progression and solo over it using notes from triads")
	}
}

func main() {

	doc := DocWriter{
		W: os.Stdout,
		Depth: 1,
	}
	

	currentTime := time.Now()
	
	doc.EmitPreamble()
	
	doc.EmitHeader("Triad Practice")
	doc.EmitContentLn("\nStats:\n")
	doc.EmitContentLn("Generated %s\n", FormatDate(currentTime))
	doc.EmitContentLn("%d major guitar keys: %s", len(MUSIC_KEYS), strings.Join(MusicKeysSlice(), " "))
	doc.EmitContentLn("%d minor guitar keys: %sm", len(MUSIC_KEYS), strings.Join(MusicKeysSlice(), "m "))

	
	numPracticeSessions := len(MUSIC_KEYS) * 2 * 4
	doc.EmitContentLn("Total practice sessions: %d", numPracticeSessions)
	
	doc.Depth++

	// emit two major keys, then the two keys in minor, repeat
	for keyIndex := 0; keyIndex < len(MUSIC_KEYS);  keyIndex += 2 {
		
		EmitKey(doc, MUSIC_KEYS[keyIndex:keyIndex+1], true)

		if keyIndex+1 < len(MUSIC_KEYS) {
			EmitKey(doc, MUSIC_KEYS[keyIndex+1:keyIndex+2], true)
		}

		EmitKey(doc, MUSIC_KEYS[keyIndex:keyIndex+1], false)

		if keyIndex+1 < len(MUSIC_KEYS) {
			EmitKey(doc, MUSIC_KEYS[keyIndex+1:keyIndex+2], false)
		}
		
	}
}
