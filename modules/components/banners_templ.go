// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type Banner struct {
	Src string
}

type BannersSchema struct {
	Banners []Banner
}

var bannersContent = BannersSchema{
	Banners: []Banner{
		{
			Src: "/static/images/banner-desktop.png",
		},
		{
			Src: "/static/images/banner-desktop.png",
		},
		{
			Src: "/static/images/banner-desktop.png",
		},
	},
}

func BannersScript() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_BannersScript_e879`,
		Function: `function __templ_BannersScript_e879(){const OPTIONS = { loop: true }
  const emblaNode = document.querySelector('#main-banners .embla')

  const prevBtnNode = document.querySelector('#main-banners .embla__button--prev')
  const nextBtnNode = document.querySelector('#main-banners .embla__button--next')
  const dotsNode = document.querySelector('#main-banners .embla__dots')

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

  const addDotBtnsAndClickHandlers = (emblaApi, dotsNode) => {
    let dotNodes = []

    const addDotBtnsWithClickHandlers = () => {
      dotsNode.innerHTML = emblaApi
        .scrollSnapList()
        .map(() => '<button class="embla__dot w-4 h-4 grid place-items-center" type="button"><div class="rounded-full duration-300 transition-all"/></button>')
        .join("")

      const scrollTo = (index) => {
        emblaApi.scrollTo(index)
      }

      dotNodes = Array.from(dotsNode.querySelectorAll('.embla__dot'))
      dotNodes.forEach((dotNode, index) => {
        dotNode.addEventListener('click', () => scrollTo(index), false)
      })
    }

    const toggleDotBtnsActive = () => {
      const selected = emblaApi.selectedScrollSnap()

      dotNodes.forEach((dotNode, index) => {
        if(index === selected){
          dotNodes[index].querySelector('div').classList.remove('w-2', 'h-2', 'bg-neutral-400')
          dotNodes[index].querySelector('div').classList.add('w-4', 'h-4', 'bg-neutral-500')
        }else{
          dotNodes[index].querySelector('div').classList.remove('w-4', 'h-4', 'bg-neutral-500')
          dotNodes[index].querySelector('div').classList.add('w-2', 'h-2', 'bg-neutral-400')
        }
      })
    }

    emblaApi
      .on('init', addDotBtnsWithClickHandlers)
      .on('reInit', addDotBtnsWithClickHandlers)
      .on('init', toggleDotBtnsActive)
      .on('reInit', toggleDotBtnsActive)
      .on('select', toggleDotBtnsActive)

    return () => {
      dotsNode.innerHTML = ""
    }
  }

  const removePrevNextBtnsClickHandlers = addPrevNextBtnsClickHandlers(
    emblaApi,
    prevBtnNode,
    nextBtnNode
  )

  const removeDotBtnsAndClickHandlers = addDotBtnsAndClickHandlers(
    emblaApi,
    dotsNode
  )

  emblaApi.on('destroy', removePrevNextBtnsClickHandlers)
  emblaApi.on('destroy', removeDotBtnsAndClickHandlers)
}`,
		Call:       templ.SafeScript(`__templ_BannersScript_e879`),
		CallInline: templ.SafeScriptInline(`__templ_BannersScript_e879`),
	}
}

func Banners() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"main-banners\" class=\"relative\"><div class=\"embla overflow-hidden\"><div class=\"embla__container flex\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, banner := range bannersContent.Banners {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"embla__slide flex-[0_0_100%] min-w-0\"><img alt=\"\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(banner.Src))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full\"></span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"p-4 absolute top-1/2 -translate-y-1/2 w-full flex justify-between items-center text-neutral-400\"><button class=\"embla__button--prev\"><img src=\"/static/svg/arrow.svg\" class=\"shrink-0\"></button> <button class=\"embla__button--next rotate-180\"><img src=\"/static/svg/arrow.svg\" class=\"shrink-0\"></button></div><div class=\"embla__dots absolute bottom-4 left-0 w-full flex justify-center gap-1\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = BannersScript().Render(ctx, templ_7745c5c3_Buffer)
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
