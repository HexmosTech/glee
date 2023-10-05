code_highlighter_theme_monokai = """
<style>
   pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .codehilite .hll { background-color: #49483e }
   .codehilite  { background: #303030; color: #f8f8f2 }
   .codehilite .c { color: #75715e } /* Comment */
   .codehilite .err { color: #960050; background-color: #1e0010 } /* Error */
   .codehilite .k { color: #66d9ef } /* Keyword */
   .codehilite .l { color: #ae81ff } /* Literal */
   .codehilite .n { color: #f8f8f2 } /* Name */
   .codehilite .o { color: #f92672 } /* Operator */
   .codehilite .p { color: #f8f8f2 } /* Punctuation */
   .codehilite .ch { color: #75715e } /* Comment.Hashbang */
   .codehilite .cm { color: #75715e } /* Comment.Multiline */
   .codehilite .cp { color: #75715e } /* Comment.Preproc */
   .codehilite .cpf { color: #75715e } /* Comment.PreprocFile */
   .codehilite .c1 { color: #75715e } /* Comment.Single */
   .codehilite .cs { color: #75715e } /* Comment.Special */
   .codehilite .gd { color: #f92672 } /* Generic.Deleted */
   .codehilite .ge { font-style: italic } /* Generic.Emph */
   .codehilite .gi { color: #a6e22e } /* Generic.Inserted */
   .codehilite .gs { font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #75715e } /* Generic.Subheading */
   .codehilite .kc { color: #66d9ef } /* Keyword.Constant */
   .codehilite .kd { color: #66d9ef } /* Keyword.Declaration */
   .codehilite .kn { color: #f92672 } /* Keyword.Namespace */
   .codehilite .kp { color: #66d9ef } /* Keyword.Pseudo */
   .codehilite .kr { color: #66d9ef } /* Keyword.Reserved */
   .codehilite .kt { color: #66d9ef } /* Keyword.Type */
   .codehilite .ld { color: #e6db74 } /* Literal.Date */
   .codehilite .m { color: #ae81ff } /* Literal.Number */
   .codehilite .s { color: #e6db74 } /* Literal.String */
   .codehilite .na { color: #a6e22e } /* Name.Attribute */
   .codehilite .nb { color: #f8f8f2 } /* Name.Builtin */
   .codehilite .nc { color: #a6e22e } /* Name.Class */
   .codehilite .no { color: #66d9ef } /* Name.Constant */
   .codehilite .nd { color: #a6e22e } /* Name.Decorator */
   .codehilite .ni { color: #f8f8f2 } /* Name.Entity */
   .codehilite .ne { color: #a6e22e } /* Name.Exception */
   .codehilite .nf { color: #a6e22e } /* Name.Function */
   .codehilite .nl { color: #f8f8f2 } /* Name.Label */
   .codehilite .nn { color: #f8f8f2 } /* Name.Namespace */
   .codehilite .nx { color: #a6e22e } /* Name.Other */
   .codehilite .py { color: #f8f8f2 } /* Name.Property */
   .codehilite .nt { color: #f92672 } /* Name.Tag */
   .codehilite .nv { color: #f8f8f2 } /* Name.Variable */
   .codehilite .ow { color: #f92672 } /* Operator.Word */
   .codehilite .w { color: #f8f8f2 } /* Text.Whitespace */
   .codehilite .mb { color: #ae81ff } /* Literal.Number.Bin */
   .codehilite .mf { color: #ae81ff } /* Literal.Number.Float */
   .codehilite .mh { color: #ae81ff } /* Literal.Number.Hex */
   .codehilite .mi { color: #ae81ff } /* Literal.Number.Integer */
   .codehilite .mo { color: #ae81ff } /* Literal.Number.Oct */
   .codehilite .sa { color: #e6db74 } /* Literal.String.Affix */
   .codehilite .sb { color: #e6db74 } /* Literal.String.Backtick */
   .codehilite .sc { color: #e6db74 } /* Literal.String.Char */
   .codehilite .dl { color: #e6db74 } /* Literal.String.Delimiter */
   .codehilite .sd { color: #e6db74 } /* Literal.String.Doc */
   .codehilite .s2 { color: #e6db74 } /* Literal.String.Double */
   .codehilite .se { color: #ae81ff } /* Literal.String.Escape */
   .codehilite .sh { color: #e6db74 } /* Literal.String.Heredoc */
   .codehilite .si { color: #e6db74 } /* Literal.String.Interpol */
   .codehilite .sx { color: #e6db74 } /* Literal.String.Other */
   .codehilite .sr { color: #e6db74 } /* Literal.String.Regex */
   .codehilite .s1 { color: #e6db74 } /* Literal.String.Single */
   .codehilite .ss { color: #e6db74 } /* Literal.String.Symbol */
   .codehilite .bp { color: #f8f8f2 } /* Name.Builtin.Pseudo */
   .codehilite .fm { color: #a6e22e } /* Name.Function.Magic */
   .codehilite .vc { color: #f8f8f2 } /* Name.Variable.Class */
   .codehilite .vg { color: #f8f8f2 } /* Name.Variable.Global */
   .codehilite .vi { color: #f8f8f2 } /* Name.Variable.Instance */
   .codehilite .vm { color: #f8f8f2 } /* Name.Variable.Magic */
   .codehilite .il { color: #ae81ff } /* Literal.Number.Integer.Long */

   .article-image {
   max-width: 600px;
   margin: 0 auto !important;
   float: none !important;
   }
</style>
"""

code_highlighter_theme_murphy = """
<style>
   pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .codehilite .hll { background-color: #ffffcc }
   .codehilite  { background: #ffffff; }
   .codehilite .c { color: #666666; font-style: italic } /* Comment */
   .codehilite .err { color: #FF0000; background-color: #FFAAAA } /* Error */
   .codehilite .k { color: #228899; font-weight: bold } /* Keyword */
   .codehilite .o { color: #333333 } /* Operator */
   .codehilite .ch { color: #666666; font-style: italic } /* Comment.Hashbang */
   .codehilite .cm { color: #666666; font-style: italic } /* Comment.Multiline */
   .codehilite .cp { color: #557799 } /* Comment.Preproc */
   .codehilite .cpf { color: #666666; font-style: italic } /* Comment.PreprocFile */
   .codehilite .c1 { color: #666666; font-style: italic } /* Comment.Single */
   .codehilite .cs { color: #cc0000; font-weight: bold; font-style: italic } /* Comment.Special */
   .codehilite .gd { color: #A00000 } /* Generic.Deleted */
   .codehilite .ge { font-style: italic } /* Generic.Emph */
   .codehilite .gr { color: #FF0000 } /* Generic.Error */
   .codehilite .gh { color: #000080; font-weight: bold } /* Generic.Heading */
   .codehilite .gi { color: #00A000 } /* Generic.Inserted */
   .codehilite .go { color: #888888 } /* Generic.Output */
   .codehilite .gp { color: #c65d09; font-weight: bold } /* Generic.Prompt */
   .codehilite .gs { font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #800080; font-weight: bold } /* Generic.Subheading */
   .codehilite .gt { color: #0044DD } /* Generic.Traceback */
   .codehilite .kc { color: #228899; font-weight: bold } /* Keyword.Constant */
   .codehilite .kd { color: #228899; font-weight: bold } /* Keyword.Declaration */
   .codehilite .kn { color: #228899; font-weight: bold } /* Keyword.Namespace */
   .codehilite .kp { color: #0088ff; font-weight: bold } /* Keyword.Pseudo */
   .codehilite .kr { color: #228899; font-weight: bold } /* Keyword.Reserved */
   .codehilite .kt { color: #6666ff; font-weight: bold } /* Keyword.Type */
   .codehilite .m { color: #6600EE; font-weight: bold } /* Literal.Number */
   .codehilite .s { background-color: #e0e0ff } /* Literal.String */
   .codehilite .na { color: #000077 } /* Name.Attribute */
   .codehilite .nb { color: #007722 } /* Name.Builtin */
   .codehilite .nc { color: #ee99ee; font-weight: bold } /* Name.Class */
   .codehilite .no { color: #55eedd; font-weight: bold } /* Name.Constant */
   .codehilite .nd { color: #555555; font-weight: bold } /* Name.Decorator */
   .codehilite .ni { color: #880000 } /* Name.Entity */
   .codehilite .ne { color: #FF0000; font-weight: bold } /* Name.Exception */
   .codehilite .nf { color: #55eedd; font-weight: bold } /* Name.Function */
   .codehilite .nl { color: #997700; font-weight: bold } /* Name.Label */
   .codehilite .nn { color: #0e84b5; font-weight: bold } /* Name.Namespace */
   .codehilite .nt { color: #007700 } /* Name.Tag */
   .codehilite .nv { color: #003366 } /* Name.Variable */
   .codehilite .ow { color: #000000; font-weight: bold } /* Operator.Word */
   .codehilite .w { color: #bbbbbb } /* Text.Whitespace */
   .codehilite .mb { color: #6600EE; font-weight: bold } /* Literal.Number.Bin */
   .codehilite .mf { color: #6600EE; font-weight: bold } /* Literal.Number.Float */
   .codehilite .mh { color: #005588; font-weight: bold } /* Literal.Number.Hex */
   .codehilite .mi { color: #6666ff; font-weight: bold } /* Literal.Number.Integer */
   .codehilite .mo { color: #4400EE; font-weight: bold } /* Literal.Number.Oct */
   .codehilite .sa { background-color: #e0e0ff } /* Literal.String.Affix */
   .codehilite .sb { background-color: #e0e0ff } /* Literal.String.Backtick */
   .codehilite .sc { color: #8888FF } /* Literal.String.Char */
   .codehilite .dl { background-color: #e0e0ff } /* Literal.String.Delimiter */
   .codehilite .sd { color: #DD4422 } /* Literal.String.Doc */
   .codehilite .s2 { background-color: #e0e0ff } /* Literal.String.Double */
   .codehilite .se { color: #666666; font-weight: bold; background-color: #e0e0ff } /* Literal.String.Escape */
   .codehilite .sh { background-color: #e0e0ff } /* Literal.String.Heredoc */
   .codehilite .si { background-color: #eeeeee } /* Literal.String.Interpol */
   .codehilite .sx { color: #ff8888; background-color: #e0e0ff } /* Literal.String.Other */
   .codehilite .sr { color: #000000; background-color: #e0e0ff } /* Literal.String.Regex */
   .codehilite .s1 { background-color: #e0e0ff } /* Literal.String.Single */
   .codehilite .ss { color: #ffcc88 } /* Literal.String.Symbol */
   .codehilite .bp { color: #007722 } /* Name.Builtin.Pseudo */
   .codehilite .fm { color: #55eedd; font-weight: bold } /* Name.Function.Magic */
   .codehilite .vc { color: #ccccff } /* Name.Variable.Class */
   .codehilite .vg { color: #ff8844 } /* Name.Variable.Global */
   .codehilite .vi { color: #aaaaff } /* Name.Variable.Instance */
   .codehilite .vm { color: #003366 } /* Name.Variable.Magic */
   .codehilite .il { color: #6666ff; font-weight: bold } /* Literal.Number.Integer.Long */
</style>"""

code_highlighter_theme_tango = """
<style>
   pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .codehilite .hll { background-color: #ffffcc }
   .codehilite  { background: #f8f8f8; }
   .codehilite .c { color: #8f5902; font-style: italic } /* Comment */
   .codehilite .err { color: #a40000; border: 1px solid #ef2929 } /* Error */
   .codehilite .g { color: #000000 } /* Generic */
   .codehilite .k { color: #204a87; font-weight: bold } /* Keyword */
   .codehilite .l { color: #000000 } /* Literal */
   .codehilite .n { color: #000000 } /* Name */
   .codehilite .o { color: #ce5c00; font-weight: bold } /* Operator */
   .codehilite .x { color: #000000 } /* Other */
   .codehilite .p { color: #000000; font-weight: bold } /* Punctuation */
   .codehilite .ch { color: #8f5902; font-style: italic } /* Comment.Hashbang */
   .codehilite .cm { color: #8f5902; font-style: italic } /* Comment.Multiline */
   .codehilite .cp { color: #8f5902; font-style: italic } /* Comment.Preproc */
   .codehilite .cpf { color: #8f5902; font-style: italic } /* Comment.PreprocFile */
   .codehilite .c1 { color: #8f5902; font-style: italic } /* Comment.Single */
   .codehilite .cs { color: #8f5902; font-style: italic } /* Comment.Special */
   .codehilite .gd { color: #a40000 } /* Generic.Deleted */
   .codehilite .ge { color: #000000; font-style: italic } /* Generic.Emph */
   .codehilite .gr { color: #ef2929 } /* Generic.Error */
   .codehilite .gh { color: #000080; font-weight: bold } /* Generic.Heading */
   .codehilite .gi { color: #00A000 } /* Generic.Inserted */
   .codehilite .go { color: #000000; font-style: italic } /* Generic.Output */
   .codehilite .gp { color: #8f5902 } /* Generic.Prompt */
   .codehilite .gs { color: #000000; font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #800080; font-weight: bold } /* Generic.Subheading */
   .codehilite .gt { color: #a40000; font-weight: bold } /* Generic.Traceback */
   .codehilite .kc { color: #204a87; font-weight: bold } /* Keyword.Constant */
   .codehilite .kd { color: #204a87; font-weight: bold } /* Keyword.Declaration */
   .codehilite .kn { color: #204a87; font-weight: bold } /* Keyword.Namespace */
   .codehilite .kp { color: #204a87; font-weight: bold } /* Keyword.Pseudo */
   .codehilite .kr { color: #204a87; font-weight: bold } /* Keyword.Reserved */
   .codehilite .kt { color: #204a87; font-weight: bold } /* Keyword.Type */
   .codehilite .ld { color: #000000 } /* Literal.Date */
   .codehilite .m { color: #0000cf; font-weight: bold } /* Literal.Number */
   .codehilite .s { color: #4e9a06 } /* Literal.String */
   .codehilite .na { color: #c4a000 } /* Name.Attribute */
   .codehilite .nb { color: #204a87 } /* Name.Builtin */
   .codehilite .nc { color: #000000 } /* Name.Class */
   .codehilite .no { color: #000000 } /* Name.Constant */
   .codehilite .nd { color: #5c35cc; font-weight: bold } /* Name.Decorator */
   .codehilite .ni { color: #ce5c00 } /* Name.Entity */
   .codehilite .ne { color: #cc0000; font-weight: bold } /* Name.Exception */
   .codehilite .nf { color: #000000 } /* Name.Function */
   .codehilite .nl { color: #f57900 } /* Name.Label */
   .codehilite .nn { color: #000000 } /* Name.Namespace */
   .codehilite .nx { color: #000000 } /* Name.Other */
   .codehilite .py { color: #000000 } /* Name.Property */
   .codehilite .nt { color: #204a87; font-weight: bold } /* Name.Tag */
   .codehilite .nv { color: #000000 } /* Name.Variable */
   .codehilite .ow { color: #204a87; font-weight: bold } /* Operator.Word */
   .codehilite .w { color: #f8f8f8; text-decoration: underline } /* Text.Whitespace */
   .codehilite .mb { color: #0000cf; font-weight: bold } /* Literal.Number.Bin */
   .codehilite .mf { color: #0000cf; font-weight: bold } /* Literal.Number.Float */
   .codehilite .mh { color: #0000cf; font-weight: bold } /* Literal.Number.Hex */
   .codehilite .mi { color: #0000cf; font-weight: bold } /* Literal.Number.Integer */
   .codehilite .mo { color: #0000cf; font-weight: bold } /* Literal.Number.Oct */
   .codehilite .sa { color: #4e9a06 } /* Literal.String.Affix */
   .codehilite .sb { color: #4e9a06 } /* Literal.String.Backtick */
   .codehilite .sc { color: #4e9a06 } /* Literal.String.Char */
   .codehilite .dl { color: #4e9a06 } /* Literal.String.Delimiter */
   .codehilite .sd { color: #8f5902; font-style: italic } /* Literal.String.Doc */
   .codehilite .s2 { color: #4e9a06 } /* Literal.String.Double */
   .codehilite .se { color: #4e9a06 } /* Literal.String.Escape */
   .codehilite .sh { color: #4e9a06 } /* Literal.String.Heredoc */
   .codehilite .si { color: #4e9a06 } /* Literal.String.Interpol */
   .codehilite .sx { color: #4e9a06 } /* Literal.String.Other */
   .codehilite .sr { color: #4e9a06 } /* Literal.String.Regex */
   .codehilite .s1 { color: #4e9a06 } /* Literal.String.Single */
   .codehilite .ss { color: #4e9a06 } /* Literal.String.Symbol */
   .codehilite .bp { color: #3465a4 } /* Name.Builtin.Pseudo */
   .codehilite .fm { color: #000000 } /* Name.Function.Magic */
   .codehilite .vc { color: #000000 } /* Name.Variable.Class */
   .codehilite .vg { color: #000000 } /* Name.Variable.Global */
   .codehilite .vi { color: #000000 } /* Name.Variable.Instance */
   .codehilite .vm { color: #000000 } /* Name.Variable.Magic */
   .codehilite .il { color: #0000cf; font-weight: bold } /* Literal.Number.Integer.Long */
</style>
"""

code_highlighter_theme_borland ="""
<style>
   pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .codehilite .hll { background-color: #ffffcc }
   .codehilite  { background: #ffffff; }
   .codehilite .c { color: #008800; font-style: italic } /* Comment */
   .codehilite .err { color: #a61717; background-color: #e3d2d2 } /* Error */
   .codehilite .k { color: #000080; font-weight: bold } /* Keyword */
   .codehilite .ch { color: #008800; font-style: italic } /* Comment.Hashbang */
   .codehilite .cm { color: #008800; font-style: italic } /* Comment.Multiline */
   .codehilite .cp { color: #008080 } /* Comment.Preproc */
   .codehilite .cpf { color: #008800; font-style: italic } /* Comment.PreprocFile */
   .codehilite .c1 { color: #008800; font-style: italic } /* Comment.Single */
   .codehilite .cs { color: #008800; font-weight: bold } /* Comment.Special */
   .codehilite .gd { color: #000000; background-color: #ffdddd } /* Generic.Deleted */
   .codehilite .ge { font-style: italic } /* Generic.Emph */
   .codehilite .gr { color: #aa0000 } /* Generic.Error */
   .codehilite .gh { color: #999999 } /* Generic.Heading */
   .codehilite .gi { color: #000000; background-color: #ddffdd } /* Generic.Inserted */
   .codehilite .go { color: #888888 } /* Generic.Output */
   .codehilite .gp { color: #555555 } /* Generic.Prompt */
   .codehilite .gs { font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #aaaaaa } /* Generic.Subheading */
   .codehilite .gt { color: #aa0000 } /* Generic.Traceback */
   .codehilite .kc { color: #000080; font-weight: bold } /* Keyword.Constant */
   .codehilite .kd { color: #000080; font-weight: bold } /* Keyword.Declaration */
   .codehilite .kn { color: #000080; font-weight: bold } /* Keyword.Namespace */
   .codehilite .kp { color: #000080; font-weight: bold } /* Keyword.Pseudo */
   .codehilite .kr { color: #000080; font-weight: bold } /* Keyword.Reserved */
   .codehilite .kt { color: #000080; font-weight: bold } /* Keyword.Type */
   .codehilite .m { color: #0000FF } /* Literal.Number */
   .codehilite .s { color: #0000FF } /* Literal.String */
   .codehilite .na { color: #FF0000 } /* Name.Attribute */
   .codehilite .nt { color: #000080; font-weight: bold } /* Name.Tag */
   .codehilite .ow { font-weight: bold } /* Operator.Word */
   .codehilite .w { color: #bbbbbb } /* Text.Whitespace */
   .codehilite .mb { color: #0000FF } /* Literal.Number.Bin */
   .codehilite .mf { color: #0000FF } /* Literal.Number.Float */
   .codehilite .mh { color: #0000FF } /* Literal.Number.Hex */
   .codehilite .mi { color: #0000FF } /* Literal.Number.Integer */
   .codehilite .mo { color: #0000FF } /* Literal.Number.Oct */
   .codehilite .sa { color: #0000FF } /* Literal.String.Affix */
   .codehilite .sb { color: #0000FF } /* Literal.String.Backtick */
   .codehilite .sc { color: #800080 } /* Literal.String.Char */
   .codehilite .dl { color: #0000FF } /* Literal.String.Delimiter */
   .codehilite .sd { color: #0000FF } /* Literal.String.Doc */
   .codehilite .s2 { color: #0000FF } /* Literal.String.Double */
   .codehilite .se { color: #0000FF } /* Literal.String.Escape */
   .codehilite .sh { color: #0000FF } /* Literal.String.Heredoc */
   .codehilite .si { color: #0000FF } /* Literal.String.Interpol */
   .codehilite .sx { color: #0000FF } /* Literal.String.Other */
   .codehilite .sr { color: #0000FF } /* Literal.String.Regex */
   .codehilite .s1 { color: #0000FF } /* Literal.String.Single */
   .codehilite .ss { color: #0000FF } /* Literal.String.Symbol */
   .codehilite .il { color: #0000FF } /* Literal.Number.Integer.Long */
</style>

"""

code_highlighter_theme_trac= """
   pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .codehilite .hll { background-color: #ffffcc }
   .codehilite  { background: #ffffff; }
   .codehilite .c { color: #999988; font-style: italic } /* Comment */
   .codehilite .err { color: #a61717; background-color: #e3d2d2 } /* Error */
   .codehilite .k { font-weight: bold } /* Keyword */
   .codehilite .o { font-weight: bold } /* Operator */
   .codehilite .ch { color: #999988; font-style: italic } /* Comment.Hashbang */
   .codehilite .cm { color: #999988; font-style: italic } /* Comment.Multiline */
   .codehilite .cp { color: #999999; font-weight: bold } /* Comment.Preproc */
   .codehilite .cpf { color: #999988; font-style: italic } /* Comment.PreprocFile */
   .codehilite .c1 { color: #999988; font-style: italic } /* Comment.Single */
   .codehilite .cs { color: #999999; font-weight: bold; font-style: italic } /* Comment.Special */
   .codehilite .gd { color: #000000; background-color: #ffdddd } /* Generic.Deleted */
   .codehilite .ge { font-style: italic } /* Generic.Emph */
   .codehilite .gr { color: #aa0000 } /* Generic.Error */
   .codehilite .gh { color: #999999 } /* Generic.Heading */
   .codehilite .gi { color: #000000; background-color: #ddffdd } /* Generic.Inserted */
   .codehilite .go { color: #888888 } /* Generic.Output */
   .codehilite .gp { color: #555555 } /* Generic.Prompt */
   .codehilite .gs { font-weight: bold } /* Generic.Strong */
   .codehilite .gu { color: #aaaaaa } /* Generic.Subheading */
   .codehilite .gt { color: #aa0000 } /* Generic.Traceback */
   .codehilite .kc { font-weight: bold } /* Keyword.Constant */
   .codehilite .kd { font-weight: bold } /* Keyword.Declaration */
   .codehilite .kn { font-weight: bold } /* Keyword.Namespace */
   .codehilite .kp { font-weight: bold } /* Keyword.Pseudo */
   .codehilite .kr { font-weight: bold } /* Keyword.Reserved */
   .codehilite .kt { color: #445588; font-weight: bold } /* Keyword.Type */
   .codehilite .m { color: #009999 } /* Literal.Number */
   .codehilite .s { color: #bb8844 } /* Literal.String */
   .codehilite .na { color: #008080 } /* Name.Attribute */
   .codehilite .nb { color: #999999 } /* Name.Builtin */
   .codehilite .nc { color: #445588; font-weight: bold } /* Name.Class */
   .codehilite .no { color: #008080 } /* Name.Constant */
   .codehilite .ni { color: #800080 } /* Name.Entity */
   .codehilite .ne { color: #990000; font-weight: bold } /* Name.Exception */
   .codehilite .nf { color: #990000; font-weight: bold } /* Name.Function */
   .codehilite .nn { color: #555555 } /* Name.Namespace */
   .codehilite .nt { color: #000080 } /* Name.Tag */
   .codehilite .nv { color: #008080 } /* Name.Variable */
   .codehilite .ow { font-weight: bold } /* Operator.Word */
   .codehilite .w { color: #bbbbbb } /* Text.Whitespace */
   .codehilite .mb { color: #009999 } /* Literal.Number.Bin */
   .codehilite .mf { color: #009999 } /* Literal.Number.Float */
   .codehilite .mh { color: #009999 } /* Literal.Number.Hex */
   .codehilite .mi { color: #009999 } /* Literal.Number.Integer */
   .codehilite .mo { color: #009999 } /* Literal.Number.Oct */
   .codehilite .sa { color: #bb8844 } /* Literal.String.Affix */
   .codehilite .sb { color: #bb8844 } /* Literal.String.Backtick */
   .codehilite .sc { color: #bb8844 } /* Literal.String.Char */
   .codehilite .dl { color: #bb8844 } /* Literal.String.Delimiter */
   .codehilite .sd { color: #bb8844 } /* Literal.String.Doc */
   .codehilite .s2 { color: #bb8844 } /* Literal.String.Double */
   .codehilite .se { color: #bb8844 } /* Literal.String.Escape */
   .codehilite .sh { color: #bb8844 } /* Literal.String.Heredoc */
   .codehilite .si { color: #bb8844 } /* Literal.String.Interpol */
   .codehilite .sx { color: #bb8844 } /* Literal.String.Other */
   .codehilite .sr { color: #808000 } /* Literal.String.Regex */
   .codehilite .s1 { color: #bb8844 } /* Literal.String.Single */
   .codehilite .ss { color: #bb8844 } /* Literal.String.Symbol */
   .codehilite .bp { color: #999999 } /* Name.Builtin.Pseudo */
   .codehilite .fm { color: #990000; font-weight: bold } /* Name.Function.Magic */
   .codehilite .vc { color: #008080 } /* Name.Variable.Class */
   .codehilite .vg { color: #008080 } /* Name.Variable.Global */
   .codehilite .vi { color: #008080 } /* Name.Variable.Instance */
   .codehilite .vm { color: #008080 } /* Name.Variable.Magic */
   .codehilite .il { color: #009999 } /* Literal.Number.Integer.Long */
"""

sidebar_toc_head = """
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.css">
<style>
   .gh-content {
   position: relative;
   }
   .gh-toc > .toc-list {
   position: relative;
   }
   .toc-list {
   overflow: hidden;
   list-style: none;
   }
   @media (min-width: 1300px) {
   .gh-sidebar {
   position: absolute; 
   top: 0;
   bottom: 0;
   margin-top: 4vmin;
   margin-left: 20px;
   grid-column: wide-end / main-end; /* Place the TOC to the right of the content */
   width: inline-block;
   white-space: nowrap;
   }
   .gh-toc-container {
   position: sticky; /* On larger screens, TOC will stay in the same spot on the page */
   top: 4vmin;
   }
   }
   .gh-toc .is-active-link::before {
   background-color: var(--ghost-accent-color); /* Defines TOC accent color based on Accent color set in Ghost Admin */
   } 
</style>
"""

sidebar_toc_footer = """
<script src="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.min.js"></script>
<script>
   const parent = document.querySelector(".gh-content.gh-canvas");
   // Create the <aside> element
   const asideElement = document.createElement("aside");
   asideElement.setAttribute("class", "gh-sidebar");
   //asideElement.style.zIndex = 0; // sent to back so it doesn't show on top of images
   
   // Create the container div for title and TOC
   const containerElement = document.createElement("div");
   containerElement.setAttribute("class", "gh-toc-container");
   
   // Create the title element
   const titleElement = document.createElement("div");
   titleElement.textContent = "Table of Contents";
   titleElement.style.fontWeight = "bold";
   containerElement.appendChild(titleElement);
   
   // Create the <div> element for TOC
   const divElement = document.createElement("div");
   divElement.setAttribute("class", "gh-toc");
   containerElement.appendChild(divElement);
   
   // Append the <div> element to the <aside> element
   asideElement.appendChild(containerElement);
   parent.insertBefore(asideElement, parent.firstChild);
   
   tocbot.init({
       // Where to render the table of contents.
       tocSelector: '.gh-toc',
       // Where to grab the headings to build the table of contents.
       contentSelector: '.gh-content',
       // Which headings to grab inside of the contentSelector element.
       headingSelector: 'h1, h2, h3, h4',
       // Ensure correct positioning
       hasInnerContainers: true,
   });
   
   // Get the table of contents element
   const toc = document.querySelector(".gh-toc");
   const sidebar = document.querySelector(".gh-sidebar");
   
   // Check the number of items in the table of contents
   const tocItems = toc.querySelectorAll('li').length;
   
   // Only show the table of contents if it has more than 5 items
   if (tocItems > 2) {
     sidebar.style.display = 'block';
   } else {
     sidebar.style.display = 'none';
   }
</script>
"""
