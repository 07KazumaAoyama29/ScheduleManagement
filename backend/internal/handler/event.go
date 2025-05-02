package handler

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "example.com/calendarapp/internal/domain/model"
)

type EventHandler struct {
    db *sql.DB
}

func NewEventHandler(db *sql.DB) *EventHandler { return &EventHandler{db: db} }

type createEventReq struct {
    Title   string    `json:"title"   binding:"required"`
    StartAt time.Time `json:"startAt" binding:"required"`
    EndAt   time.Time `json:"endAt"   binding:"required"`
    RRule   string    `json:"rrule"`
}

func (h *EventHandler) List(c *gin.Context) {
    rows, err := h.db.Query(`SELECT id,title,start_at,end_at,rrule FROM events ORDER BY start_at`)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var events []model.Event
    for rows.Next() {
        var ev model.Event
        if err := rows.Scan(&ev.ID, &ev.Title, &ev.StartAt, &ev.EndAt, &ev.RRule); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        events = append(events, ev)
    }
    c.JSON(http.StatusOK, events)
}

func (h *EventHandler) Create(c *gin.Context) {
    var req createEventReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id := uuid.NewString()
    _, err := h.db.Exec(`INSERT INTO events(id,title,start_at,end_at,rrule) VALUES($1,$2,$3,$4,$5)`,
        id, req.Title, req.StartAt, req.EndAt, req.RRule)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": id})
}