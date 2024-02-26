// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type Slide struct {
	Label     string
	LinkLabel string
	Href      string
}

type TopbarSchema struct {
	slides []Slide
}

var topbarContent = TopbarSchema{
	slides: []Slide{
		{
			Label:     "Frete Grátis acima de R$00,00",
			LinkLabel: "saiba mais",
			Href:      "#",
		},
		{
			Label:     "Frete Grátis acima de R$00,00",
			LinkLabel: "saiba mais",
			Href:      "#",
		},
		{
			Label:     "Frete Grátis acima de R$00,00",
			LinkLabel: "saiba mais",
			Href:      "#",
		},
	},
}

func TopbarScript() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_TopbarScript_9e20`,
		Function: `function __templ_TopbarScript_9e20(){const OPTIONS = { loop: true }
  const emblaNode = document.querySelector('#topbar .embla')

  const prevBtnNode = document.querySelector('#topbar .embla__button--prev')
  const nextBtnNode = document.querySelector('#topbar .embla__button--next')

  const emblaApi = EmblaCarousel(emblaNode, OPTIONS)

  const addTogglePrevNextBtnsActive = (emblaApi, prevBtn, nextBtn) => {
    const togglePrevNextBtnsState = () => {
      if (emblaApi.canScrollPrev()) prevBtn.removeAttribute('disabled')
      else prevBtn.setAttribute('disabled', 'disabled')

      if (emblaApi.canScrollNext()) nextBtn.removeAttribute('disabled')
      else nextBtn.setAttribute('disabled', 'disabled')
    }

    emblaApi
      .on('select', togglePrevNextBtnsState)
      .on('init', togglePrevNextBtnsState)
      .on('reInit', togglePrevNextBtnsState)

    return () => {
      prevBtn.removeAttribute('disabled')
      nextBtn.removeAttribute('disabled')
    }
  }

  const addPrevNextBtnsClickHandlers = (emblaApi, prevBtn, nextBtn) => {
    const scrollPrev = () => {
      emblaApi.scrollPrev()
    }
    const scrollNext = () => {
      emblaApi.scrollNext()
    }
    prevBtn.addEventListener('click', scrollPrev, false)
    nextBtn.addEventListener('click', scrollNext, false)

    const removeTogglePrevNextBtnsActive = addTogglePrevNextBtnsActive(
      emblaApi,
      prevBtn,
      nextBtn
    )

    return () => {
      removeTogglePrevNextBtnsActive()
      prevBtn.removeEventListener('click', scrollPrev, false)
      nextBtn.removeEventListener('click', scrollNext, false)
    }
  }

  const removePrevNextBtnsClickHandlers = addPrevNextBtnsClickHandlers(
    emblaApi,
    prevBtnNode,
    nextBtnNode
  )

  emblaApi.on('destroy', removePrevNextBtnsClickHandlers)
}`,
		Call:       templ.SafeScript(`__templ_TopbarScript_9e20`),
		CallInline: templ.SafeScriptInline(`__templ_TopbarScript_9e20`),
	}
}

func Topbar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"topbar\" class=\"bg-neutral-900 p-3 flex items-center justify-center gap-4 text-neutral-100\"><button class=\"embla__button--prev\"><img src=\"/static/svg/chevron.svg\" class=\"shrink-0\"></button><div class=\"embla overflow-hidden\"><div class=\"embla__container flex\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range topbarContent.slides {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"embla__slide text-sm text-center flex-[0_0_100%] min-w-0\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(item.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules\components\topbar.templ`, Line: 102, Col: 18}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(item.LinkLabel) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 templ.SafeURL = templ.URL(item.Href)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-neutral-100/60\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(item.LinkLabel)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `modules\components\topbar.templ`, Line: 104, Col: 84}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><button class=\"embla__button--next rotate-180\"><img src=\"/static/svg/chevron.svg\" class=\"shrink-0\"></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TopbarScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
