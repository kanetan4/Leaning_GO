namespace go notes

struct Note {
    1: i64 id,
    2: string user
    3: string title,
    4: string content,
    5: i64 create_time
}

struct CreateNoteRequest {
    1: string user
    2: string title,
    3: string content,
}

struct GetNoteRequest {
    1: string user
    2: i64 id
}

struct DeleteNoteRequest {
    1: string user
    2: i64 id
}

service NoteService {
    Note CreateNote(1: CreateNoteRequest request)
    Note GetNote(1: GetNoteRequest request)
    list<Note> ListNotes(1: string user)
    bool DeleteNote(1: DeleteNoteRequest request)
}