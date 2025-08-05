package main

import (
	"context"
	notes "ex4_noteservice/kitex_gen/notes"
	"os"
	"time"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}


// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, request *notes.CreateNoteRequest) (resp *notes.Note, err error) {
	// TODO: Your code here...
	notesList, err := LoadNotes()
	if err != nil {
		if !os.IsNotExist(err) { //file not exist error
			return nil, err
		} else {
			notesList = []*notes.Note{}
		}
	}

	note := &notes.Note{
		Id: time.Now().UnixNano(),
		User: request.User,
		Title: request.Title,
		Content: request.Content,
		CreateTime: time.Now().UnixNano(),
	}

	notesList = append(notesList, note)

	err = SaveNotes(notesList)
	if err != nil {
		return nil, err
	}

	return note, err
}

// GetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) GetNote(ctx context.Context, request *notes.GetNoteRequest) (resp *notes.Note, err error) {
	// TODO: Your code here...
	return
}

// ListNotes implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) ListNotes(ctx context.Context, user string) (resp []*notes.Note, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, request *notes.DeleteNoteRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
