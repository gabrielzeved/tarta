// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Header() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"px-12 py-4 flex justify-between border border-neutral-200\"><div class=\"flex-1 flex items-center\"><img alt=\"logo\" class=\"w-auto h-[48px]\" src=\"/static/images/logo.png\"></div><div class=\"flex-1\"><div class=\"bg-white flex flex-row rounded overflow-hidden\"><input class=\"rounded-l w-full py-4 px-3 border border-neutral-300 focus-visible:outline-none\" type=\"text\" placeholder=\"O que você está procurando hoje?\"> <button class=\"bg-neutral-900 px-4\"><img alt=\"\" src=\"/static/svg/search.svg\"></button></div></div><div class=\"flex-1 flex items-center justify-end gap-7\"><div class=\"inline-flex items-center text-neutral-500 text-sm gap-2\"><img src=\"/static/svg/whatsapp.svg\"> <span>Compre pelo <br>Whatsapp </span></div><div class=\"inline-flex items-center text-neutral-500 text-sm gap-2\"><img src=\"/static/svg/user-circle.svg\"> <span>Entre <span class=\"text-neutral-400\">ou</span><br>Cadastre-se\r</span></div><div class=\"relative inline-block\"><img src=\"/static/svg/shopping-cart.svg\"> <span class=\"w-5 h-5 rounded-full bg-secondary-500 grid place-items-center top-0 right-0 translate-x-1/2 absolute text-xs font-bold leading-[0] text-neutral-100\">3\r</span></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}