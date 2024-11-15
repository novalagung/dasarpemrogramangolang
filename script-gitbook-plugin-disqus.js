require([
    "gitbook",
    "jQuery"
], function(gitbook, $) {
    var useIdentifier = false;
    var disqusConfig = null;

    function prepareDisqusThreadDOM() {
        var id = "disqus_thread";
        if ($("#" + id).children().length > 0) {
            return;
        }

        $("#" + id).remove()

        var $disqusDiv = $("<div>", { "id": id });
        $(".book-body .page-inner").append($disqusDiv);
    }

    function resetDisqus() {
        prepareDisqusThreadDOM()
        if (typeof DISQUS !== "undefined") {
            DISQUS.reset({
                reload: true,
                config: function () {
                    this.language = $('html').attr('lang') || "en";
                    this.page.url = window.location.href;

                    if (useIdentifier) {
                        this.page.identifier = currentUrl();
                    }
                }
            });
        }
    }

    function joinURL(baseUrl, url) {
        var theUrl = new URI(url);
        if (theUrl.is("relative")) {
            theUrl = theUrl.absoluteTo(baseUrl);
        }
        return theUrl.toString();
    }

    function currentUrl() {
        var location = new URI(window.location.href),
            base     = joinURL(window.location.href, gitbook.state.basePath),
            current  = location.relativeTo(base).toString(),
            language = $('html').attr('lang'),
            parent   = joinURL(base, '..'),
            folder   = new URI(base).relativeTo(parent).toString();

        // If parent folder is the same as language, we assume translated books
        if (folder.replace(/\/$/, "") === language) {
            current = folder + current;
        }

        return current;
    }

    function loadDisqus(config) {
        config.disqus = config.disqus || {};
        var disqus_shortname = config.disqus.shortName;
        var disqus_config = function() {
            this.language = $('html').attr('lang') || "en";
        };

        if (config.disqus.useIdentifier) {
            useIdentifier = true;
            var disqus_identifier = currentUrl();
        }

        /* * * DON'T EDIT BELOW THIS LINE * * */
        (function() {
            var dsq = document.createElement('script'); dsq.type = 'text/javascript'; dsq.async = true;
            dsq.src = '//' + disqus_shortname + '.disqus.com/embed.js';
            (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
        })();

        resetDisqus();
    }

    function lazyLoadDisqus(config) {
        prepareDisqusThreadDOM()

        if (IntersectionObserver) {
            var observer = new IntersectionObserver(function (entries) {
                if (!entries[0]) {
                    return;
                }
                if (!entries[0].isIntersecting) {
                    return;
                }
    
                // comments section reached, start loading Disqus now
                loadDisqus(config);
                observer.disconnect();
            }, {
                threshold: 0.001,
            });
            observer.observe($("#disqus_thread")[0]);
        } else {
            loadDisqus(config);
        }
    }

    gitbook.events.bind("start", function(e, config) {
        disqusConfig = config

        // lazy load disqus if possible
        lazyLoadDisqus(config)
    });

    gitbook.events.bind("page.change", function () {
        if ($("#disqus_thread").children().length > 0) {
            resetDisqus()
        } else {
            lazyLoadDisqus(disqusConfig);
        }
    });
});
