package main

func addBlogConfigurations(meta map[string]interface{}) map[string]interface{} {

	if config == nil {
		log.Fatal("Configuration is not initialized. Call loadGlobalConfig to initialize it.")
	}
   handleExcerpt(meta)

	globalSidebarTOC := config.GetDefault("blog-configuration.SIDEBAR_TOC", "").(bool)
	globalFeatured := config.GetDefault("blog-configuration.FEATURED", "").(bool)
	globalStatus := config.GetDefault("blog-configuration.STATUS", "").(string)
	if _, ok := meta["featured"]; !ok {
		meta["featured"] = globalFeatured
	}
	if _, ok := meta["status"]; !ok {
		meta["status"] = globalStatus
	}
	if meta["status"] == nil || meta["featured"] == nil {
		log.Error("required featured and status")
	}

	defaultStyle := `pre { line-height: 125%; }
   td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
   td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
   .article-image {
   max-width: 600px;
   margin: 0 auto !important;
   float: none !important;
   }`

	sidebarTocHead := `<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.css">
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
</style>`
	sidebarTocFooter := `<script src="https://cdnjs.cloudflare.com/ajax/libs/tocbot/4.12.3/tocbot.min.js"></script>
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
</script>`

	if existingHead, ok := meta["codeinjection_head"].(string); ok {
		meta["codeinjection_head"] = existingHead + "<style>" + defaultStyle + "</style>"
	} else {
		meta["codeinjection_head"] = "<style> " + defaultStyle + "</style>"
	}
	blogMetaSidebarTOC := meta["sidebar_toc"]

	if blogMetaSidebarTOC != nil {
		globalSidebarTOC = blogMetaSidebarTOC.(bool)
	}
	if globalSidebarTOC {
		if existingHead, ok := meta["codeinjection_head"].(string); ok {
			meta["codeinjection_head"] = existingHead + sidebarTocHead
			log.Debug("Done Sidebar TOC Code injection")
		} else {
			meta["codeinjection_head"] = sidebarTocHead
		}
		meta["codeinjection_foot"] = sidebarTocFooter
	}

	return meta
}

func handleExcerpt(meta map[string]interface{}) {
   // handle excerpt in the blog post.
    if _, exists := meta["custom_excerpt"]; exists {
        return
    }
    if val, ok := meta["excerpt"].(string); ok {
        meta["custom_excerpt"] = val
    } else {
        meta["custom_excerpt"] = ""
    }
}



