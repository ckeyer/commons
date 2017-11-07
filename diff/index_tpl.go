package diff
const (
	compareHTML = `<!DOCTYPE html>
<html>
  <head>
    <style>{{ .CSS }}</style>
  </head>
  <body>
    <div id="app">
      <div id="diff">
        <div class="diff-contents">
          {{ .content }}
        </div>
      </div>
    </div>
  </body>
</html>`

	cssText = `
html {
	font-family:sans-serif;
	-ms-text-size-adjust:100%;
	-webkit-text-size-adjust:100%
}
body {
	margin:0
}
article,aside,details,figcaption,figure,footer,header,hgroup,main,menu,nav,section,summary {
	display:block
}
audio,canvas,progress,video {
	display:inline-block;
	vertical-align:baseline
}
audio:not([controls]) {
	display:none;
	height:0
}
[hidden],template {
	display:none
}
a {
	background-color:transparent
}
a:active,a:hover {
	outline:0
}
abbr[title] {
	border-bottom:1px dotted
}
b,strong {
	font-weight:bold
}
dfn {
	font-style:italic
}
h1 {
	font-size:2em;
	margin:0.67em 0
}
mark {
	background:#ff0;
	color:#000
}
small {
	font-size:80%
}
sub,sup {
	font-size:75%;
	line-height:0;
	position:relative;
	vertical-align:baseline
}
sup {
	top:-0.5em
}
sub {
	bottom:-0.25em
}
img {
	border:0
}
svg:not(:root) {
	overflow:hidden
}
figure {
	margin:1em 40px
}
hr {
	box-sizing:content-box;
	height:0
}
pre {
	overflow:auto
}
code,kbd,pre,samp {
	font-family:monospace, monospace;
	font-size:1em
}
button,input,optgroup,select,textarea {
	color:inherit;
	font:inherit;
	margin:0
}
button {
	overflow:visible
}
button,select {
	text-transform:none
}
button,html input[type="button"],input[type="reset"],input[type="submit"] {
	-webkit-appearance:button;
	cursor:pointer
}
button[disabled],html input[disabled] {
	cursor:default
}
button::-moz-focus-inner,input::-moz-focus-inner {
	border:0;
	padding:0
}
input {
	line-height:normal
}
input[type="checkbox"],input[type="radio"] {
	box-sizing:border-box;
	padding:0
}
input[type="number"]::-webkit-inner-spin-button,input[type="number"]::-webkit-outer-spin-button {
	height:auto
}
input[type="search"] {
	-webkit-appearance:textfield;
	box-sizing:content-box
}
input[type="search"]::-webkit-search-cancel-button,input[type="search"]::-webkit-search-decoration {
	-webkit-appearance:none
}
fieldset {
	border:1px solid #c0c0c0;
	margin:0 2px;
	padding:0.35em 0.625em 0.75em
}
legend {
	border:0;
	padding:0
}
textarea {
	overflow:auto
}
optgroup {
	font-weight:bold
}
table {
	border-collapse:collapse;
	border-spacing:0
}
td,th {
	padding:0
}
html,body {
	height:100%;
	margin:0;
	font-family:"Source Sans Pro", "Helvetica", sans-serif;
	background:#182730
}
html {
	box-sizing:border-box
}
*,*:before,*:after {
	box-sizing:inherit
}
.ui-shadow,
#app #sidebar,
#app #sidebar .sidebar-inner .sidebar-entry:hover,
#app #sidebar .sidebar-inner .sidebar-entry.sidebar-entry-selected,
#app #diff .diff-section.diff-section-headers {
	transition:all 0.3s ease-out;
	box-shadow:0 2px 5px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12)
}
.ui-shadow-hover:hover {
	background:#26404f;
	box-shadow:0 5px 11px 0 rgba(0,0,0,0.18),0 4px 15px 0 rgba(0,0,0,0.15)
}
.hidden {
	display:none
}
#app {
	display:flex;
	min-width:100%;
	min-height:100%
}
#app #contents {
	display:flex;
	min-width:100%
}
#app .button {
	-webkit-user-select:none;
	-moz-user-select:none;
	-ms-user-select:none;
	user-select:none;
	position:absolute;
	bottom:10px;
	right:10px;
	font-family:"Helvetica Neue";
	color:white;
	background:#333;
	padding:8px 16px;
	font-size:14px;
	display:inline-block;
	border-radius:4px;
	cursor:default
}
#app .button:hover {
	background:#444
}
#app pre code {
	line-height:1.6;
	font-size:11px;
	font-family:"Menlo", "Monaco", monospace
}
#app #diff {
	-webkit-box-flex:1;
	-moz-box-flex:1;
	box-flex:1;
	-webkit-flex:1 0;
	-moz-flex:1 0;
	-ms-flex:1 0;
	flex:1 0;
	display:flex;
	flex-direction:column;
	background:white;
	overflow:hidden
}
#app #diff .diff-section {
	display:flex;
	flex-grow:1;
	min-height:100%;
	width:100%;
	padding:0;
	font-size:11px;
	font-family:"Menlo", "Monaco", monospace
}
#app #diff .diff-section.diff-section-headers {
	position:fixed;
	z-index:3;
	display:block;
	color:white;
	min-height:inherit;
	height:40px;
	background:#1E323E;
	padding:12px 10px;
	font-size:13px
}
#app #diff .diff-context-1 .diff-pane .lc-1,
#app #diff .diff-context-1 .gutter .lc-1 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-1 .diff-pane .lc-0+.lc-1,
#app #diff .diff-context-1 .gutter .lc-0+.lc-1 {
	display:block !important
}
#app #diff .diff-context-1 .diff-pane .lc-1 ~ .lc-1,
#app #diff .diff-context-1 .gutter .lc-1 ~ .lc-1 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc--1,
#app #diff .diff-context-1 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-2,
#app #diff .diff-context-1 .gutter .lc-2 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-3,
#app #diff .diff-context-1 .gutter .lc-3 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-4,
#app #diff .diff-context-1 .gutter .lc-4 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-5,
#app #diff .diff-context-1 .gutter .lc-5 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-6,
#app #diff .diff-context-1 .gutter .lc-6 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-7,
#app #diff .diff-context-1 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-8,
#app #diff .diff-context-1 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-9,
#app #diff .diff-context-1 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-1 .diff-pane .lc-10,
#app #diff .diff-context-1 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-2,
#app #diff .diff-context-2 .gutter .lc-2 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-2 .diff-pane .lc-1+.lc-2,
#app #diff .diff-context-2 .gutter .lc-1+.lc-2 {
	display:block !important
}
#app #diff .diff-context-2 .diff-pane .lc-2 ~ .lc-2,
#app #diff .diff-context-2 .gutter .lc-2 ~ .lc-2 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc--1,
#app #diff .diff-context-2 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-3,
#app #diff .diff-context-2 .gutter .lc-3 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-4,
#app #diff .diff-context-2 .gutter .lc-4 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-5,
#app #diff .diff-context-2 .gutter .lc-5 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-6,
#app #diff .diff-context-2 .gutter .lc-6 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-7,
#app #diff .diff-context-2 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-8,
#app #diff .diff-context-2 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-9,
#app #diff .diff-context-2 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-2 .diff-pane .lc-10,
#app #diff .diff-context-2 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-3,
#app #diff .diff-context-3 .gutter .lc-3 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-3 .diff-pane .lc-2+.lc-3,
#app #diff .diff-context-3 .gutter .lc-2+.lc-3 {
	display:block !important
}
#app #diff .diff-context-3 .diff-pane .lc-3 ~ .lc-3,
#app #diff .diff-context-3 .gutter .lc-3 ~ .lc-3 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc--1,
#app #diff .diff-context-3 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-4,
#app #diff .diff-context-3 .gutter .lc-4 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-5,
#app #diff .diff-context-3 .gutter .lc-5 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-6,
#app #diff .diff-context-3 .gutter .lc-6 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-7,
#app #diff .diff-context-3 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-8,
#app #diff .diff-context-3 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-9,
#app #diff .diff-context-3 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-3 .diff-pane .lc-10,
#app #diff .diff-context-3 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-4,
#app #diff .diff-context-4 .gutter .lc-4 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-4 .diff-pane .lc-3+.lc-4,
#app #diff .diff-context-4 .gutter .lc-3+.lc-4 {
	display:block !important
}
#app #diff .diff-context-4 .diff-pane .lc-4 ~ .lc-4,
#app #diff .diff-context-4 .gutter .lc-4 ~ .lc-4 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc--1,
#app #diff .diff-context-4 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-5,
#app #diff .diff-context-4 .gutter .lc-5 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-6,
#app #diff .diff-context-4 .gutter .lc-6 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-7,
#app #diff .diff-context-4 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-8,
#app #diff .diff-context-4 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-9,
#app #diff .diff-context-4 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-4 .diff-pane .lc-10,
#app #diff .diff-context-4 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-5,
#app #diff .diff-context-5 .gutter .lc-5 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-5 .diff-pane .lc-4+.lc-5,
#app #diff .diff-context-5 .gutter .lc-4+.lc-5 {
	display:block !important
}
#app #diff .diff-context-5 .diff-pane .lc-5 ~ .lc-5,
#app #diff .diff-context-5 .gutter .lc-5 ~ .lc-5 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc--1,
#app #diff .diff-context-5 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-6,
#app #diff .diff-context-5 .gutter .lc-6 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-7,
#app #diff .diff-context-5 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-8,
#app #diff .diff-context-5 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-9,
#app #diff .diff-context-5 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-5 .diff-pane .lc-10,
#app #diff .diff-context-5 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc-6,
#app #diff .diff-context-6 .gutter .lc-6 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-6 .diff-pane .lc-5+.lc-6,
#app #diff .diff-context-6 .gutter .lc-5+.lc-6 {
	display:block !important
}
#app #diff .diff-context-6 .diff-pane .lc-6 ~ .lc-6,
#app #diff .diff-context-6 .gutter .lc-6 ~ .lc-6 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc--1,
#app #diff .diff-context-6 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc-7,
#app #diff .diff-context-6 .gutter .lc-7 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc-8,
#app #diff .diff-context-6 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc-9,
#app #diff .diff-context-6 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-6 .diff-pane .lc-10,
#app #diff .diff-context-6 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-7 .diff-pane .lc-7,
#app #diff .diff-context-7 .gutter .lc-7 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-7 .diff-pane .lc-6+.lc-7,
#app #diff .diff-context-7 .gutter .lc-6+.lc-7 {
	display:block !important
}
#app #diff .diff-context-7 .diff-pane .lc-7 ~ .lc-7,
#app #diff .diff-context-7 .gutter .lc-7 ~ .lc-7 {
	display:none
}
#app #diff .diff-context-7 .diff-pane .lc--1,
#app #diff .diff-context-7 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-7 .diff-pane .lc-8,
#app #diff .diff-context-7 .gutter .lc-8 {
	display:none
}
#app #diff .diff-context-7 .diff-pane .lc-9,
#app #diff .diff-context-7 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-7 .diff-pane .lc-10,
#app #diff .diff-context-7 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-8 .diff-pane .lc-8,
#app #diff .diff-context-8 .gutter .lc-8 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-8 .diff-pane .lc-7+.lc-8,
#app #diff .diff-context-8 .gutter .lc-7+.lc-8 {
	display:block !important
}
#app #diff .diff-context-8 .diff-pane .lc-8 ~ .lc-8,
#app #diff .diff-context-8 .gutter .lc-8 ~ .lc-8 {
	display:none
}
#app #diff .diff-context-8 .diff-pane .lc--1,
#app #diff .diff-context-8 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-8 .diff-pane .lc-9,
#app #diff .diff-context-8 .gutter .lc-9 {
	display:none
}
#app #diff .diff-context-8 .diff-pane .lc-10,
#app #diff .diff-context-8 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-9 .diff-pane .lc-9,
#app #diff .diff-context-9 .gutter .lc-9 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-9 .diff-pane .lc-8+.lc-9,
#app #diff .diff-context-9 .gutter .lc-8+.lc-9 {
	display:block !important
}
#app #diff .diff-context-9 .diff-pane .lc-9 ~ .lc-9,
#app #diff .diff-context-9 .gutter .lc-9 ~ .lc-9 {
	display:none
}
#app #diff .diff-context-9 .diff-pane .lc--1,
#app #diff .diff-context-9 .gutter .lc--1 {
	display:none
}
#app #diff .diff-context-9 .diff-pane .lc-10,
#app #diff .diff-context-9 .gutter .lc-10 {
	display:none
}
#app #diff .diff-context-10 .diff-pane .lc-10,
#app #diff .diff-context-10 .gutter .lc-10 {
	height:0px !important;
	margin:14px 0;
	overflow:hidden;
	box-shadow:0 0 10px 1px rgba(0,0,0,0.2);
	opacity:0.3;
	display:block
}
#app #diff .diff-context-10 .diff-pane .lc-9+.lc-10,
#app #diff .diff-context-10 .gutter .lc-9+.lc-10 {
	display:block !important
}
#app #diff .diff-context-10 .diff-pane .lc-10 ~ .lc-10,
#app #diff .diff-context-10 .gutter .lc-10 ~ .lc-10 {
	display:none
}
#app #diff .diff-context-10 .diff-pane .lc--1,
#app #diff .diff-context-10 .gutter .lc--1 {
	display:none
}
#app #diff .diff-contents {
	display:flex;
	width:100%;
	margin-top:50px;
	border-top:1px solid #CCC
}
#app #diff .diff-contents .hljs {
	padding:0
}
#app #diff .diff-contents .gutter {
	-webkit-user-select:none;
	-moz-user-select:none;
	-ms-user-select:none;
	user-select:none;
	-webkit-box-flex:0;
	-moz-box-flex:0;
	box-flex:0;
	-webkit-flex:0 0 auto;
	-moz-flex:0 0 auto;
	-ms-flex:0 0 auto;
	flex:0 0 auto;
	font-size:10px;
	background:#fafafa;
	text-align:right;
	border-right:1px solid #CCC;
	border-left:1px solid #CCC;
	min-width:30px;
	color:rgba(0,0,0,0.4);
	z-index:2
}
#app #diff .diff-contents .gutter:first-child {
	border-left:none
}
#app #diff .diff-contents .gutter .line.line-ws,
#app #diff .diff-contents .gutter .line.ln {
	background-color:#ffffdf;
	border-left:1px solid #e6e6c2;
	border-right:1px solid #e6e6c2;
	margin-left:-1px;
	margin-right:-1px
}
#app #diff .diff-contents .line {
	height:16px;
	line-height:16px;
	padding:0px 8px
}
#app #diff .diff-contents .line:last-child {
	border-bottom:1px solid #ddd
}
#app #diff .diff-contents .line.line-ws,
#app #diff .diff-contents .line.ln,
#app #diff .diff-contents .line.la,
#app #diff .diff-contents .line.lm {
	background:white
}
#app #diff .diff-contents.diff-empty-false .line {
	display:none
}
#app #diff .diff-contents.diff-empty-false .line.line-ws,
#app #diff .diff-contents.diff-empty-false .line.ln,
#app #diff .diff-contents.diff-empty-false .line.la,
#app #diff .diff-contents.diff-empty-false .line.lm {
	display:block
}
#app #diff .diff-contents .diff-pane {
	-webkit-box-flex:1;
	-moz-box-flex:1;
	box-flex:1;
	-webkit-flex:1 1 50%;
	-moz-flex:1 1 50%;
	-ms-flex:1 1 50%;
	flex:1 1 50%;
	overflow:hidden;
	overflow-x:scroll;
	background:#eee
}
#app #diff .diff-contents .diff-pane .diff-pane-contents {
	float:left;
	background:white;
	min-width:100%
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line {
	white-space:pre;
	tab-size:4;
	background:#eee
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-add,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-del,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-edit {
	display:inline-block;
	vertical-align:bottom;
	height:16px
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-add:empty,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-del:empty,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-edit:empty {
	border-left:1px solid white
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-add:empty+span:empty,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-del:empty+span:empty,
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-edit:empty+span:empty {
	display:none
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-add {
	background:#ECFFE8
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-add:empty {
	border-color:#DBF1D7
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-del {
	background:#FFE1E2
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-del:empty {
	border-color:#ff7b7f
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line .w-edit {
	background:#ffff71
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .lm {
	opacity:0.4;
	background:white
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .line-ws {
	background-color:#ffffeb
}
#app #diff .diff-contents .diff-pane .diff-pane-contents .ln {
	background-color:#FFFFD7
}
#app #diff .diff-contents #diff-left .la {
	background-color:#FFE1E2
}
#app #diff .diff-contents #diff-right .la {
	background-color:#ECFFE8
}
#app #diff .diff-contents #gutter-left .la {
	background-color:#fffbfb;
	border-left:1px solid #ffaeb1;
	margin-left:-1px;
	margin-right:-1px;
	border-right:1px solid #ffaeb1
}
#app #diff .diff-contents #gutter-right .la {
	background-color:#DBF1D7;
	border-left:1px solid #b9e4b1;
	margin-left:-1px;
	margin-right:-1px;
	border-right:1px solid #b9e4b1
}`
)
