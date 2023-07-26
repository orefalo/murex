import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as d,c as r,d as e,b as o,w as n,e as t,f as i}from"./app-45f7c304.js";const c={},h=i(`<h1 id="murex-docs" tabindex="-1"><a class="header-anchor" href="#murex-docs" aria-hidden="true">#</a> <code>murex-docs</code></h1><blockquote><p>Displays the man pages for Murex builtins</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Displays the man pages for Murex builtins.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>murex-docs: [ flag ] command -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code># Output this man page
murex-docs: murex-docs
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--summary</code> Returns an abridged description of the command rather than the entire help page.</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>These man pages are compiled into the Murex executable.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>murex-docs</code></li><li><code>help</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),u=e("code",null,"(",-1),m=e("code",null,">>",-1),f=e("code",null,">",-1),p=e("code",null,"cast",-1),_=e("code",null,"err",-1),g=e("code",null,"man-get-flags",-1),x=e("code",null,"out",-1),b=e("code",null,"tout",-1),y=e("code",null,"tread",-1),D=e("code",null,"read",-1),T=e("em",null,"typed",-1);function S(k,w){const a=l("RouterLink");return d(),r("div",null,[h,e("ul",null,[e("li",null,[o(a,{to:"/commands/brace-quote.html"},{default:n(()=>[u,t(" (brace quote)")]),_:1}),t(": Write a string to the STDOUT without new line")]),e("li",null,[o(a,{to:"/commands/greater-than-greater-than.html"},{default:n(()=>[m,t(" (append file)")]),_:1}),t(": Writes STDIN to disk - appending contents if file already exists")]),e("li",null,[o(a,{to:"/commands/greater-than.html"},{default:n(()=>[f,t(" (truncate file)")]),_:1}),t(": Writes STDIN to disk - overwriting contents if file already exists")]),e("li",null,[o(a,{to:"/commands/cast.html"},{default:n(()=>[p]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[o(a,{to:"/commands/err.html"},{default:n(()=>[_]),_:1}),t(": Print a line to the STDERR")]),e("li",null,[o(a,{to:"/commands/man-get-flags.html"},{default:n(()=>[g]),_:1}),t(": Parses man page files for command line flags")]),e("li",null,[o(a,{to:"/commands/out.html"},{default:n(()=>[x]),_:1}),t(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[o(a,{to:"/commands/tout.html"},{default:n(()=>[b]),_:1}),t(": Print a string to the STDOUT and set it's data-type")]),e("li",null,[o(a,{to:"/commands/tread.html"},{default:n(()=>[y]),_:1}),t(": "),D,t(" a line of input from the user and store as a user defined "),T,t(" variable (deprecated)")])])])}const R=s(c,[["render",S],["__file","murex-docs.html.vue"]]);export{R as default};
