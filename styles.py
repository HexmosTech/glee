style = """
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
