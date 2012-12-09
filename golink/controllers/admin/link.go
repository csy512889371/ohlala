package admin

import (
    "github.com/QLeelulu/goku"
    // "github.com/QLeelulu/ohlala/golink"
    // "github.com/QLeelulu/ohlala/golink/filters"
    "github.com/QLeelulu/ohlala/golink/models"
    "strconv"
)

var _ = adminController.
    // index
    Get("links", admin_links).
	// 删除link
	Get("links", admin_del_links)

//

func admin_links(ctx *goku.HttpContext) goku.ActionResulter {
    links, err := models.Link_GetByPage(1, 20, "")
    if err != nil {
        ctx.ViewData["errorMsg"] = err.Error()
        return ctx.Render("error", nil)
    }
    ctx.ViewData["LinkList"] = links
    return ctx.View(nil)
}

// 删除link
func admin_del_links(ctx *goku.HttpContext) goku.ActionResulter {
	var errs string
	var ok = false

    linkId, err := strconv.ParseInt(ctx.RouteData.Params["id"], 10, 64)
	if err == nil {
		err = models.Link_DelById(linkId)
	}

    if err != nil {
        errs = err.Error()
    } else {
		ok = true
	}
    r := map[string]interface{}{
        "success": ok,
        "errors":  errs,
    }

    return ctx.Json(r)
}
