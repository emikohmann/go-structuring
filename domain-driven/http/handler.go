package http

import (
    "github.com/emikohmann/go-structuring/domain-driven/domain/ban"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Handler struct {
    serv ban.Service
}

func NewHandler(serv ban.Service) *Handler {
    return &Handler{
        serv: serv,
    }
}

func (h *Handler) BanFromReq(r *http.Request) ban.Ban {
    return ban.Ban{

    }
}

func (h *Handler) ApplyBAN(c *gin.Context) {
    if err := h.serv.Apply(h.BanFromReq(c.Request)); err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.String(http.StatusCreated, "applied")
}

func (h *Handler) RemoveBAN(c *gin.Context) {
    if err := h.serv.Remove(h.BanFromReq(c.Request)); err != nil {
        c.JSON(http.StatusInternalServerError, err)
        return
    }
    c.String(http.StatusOK, "removed")
}
