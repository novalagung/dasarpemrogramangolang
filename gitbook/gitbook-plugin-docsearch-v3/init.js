require(["gitbook"], function(gitbook) {
    let pluginsConfig = {};
    const initDocSearch = function() {
        const cfg = pluginsConfig.docsearch;
        if (!cfg.appId) {
            throw new Error('missing DocSearch appId');
        }
        if (!cfg.apiKey) {
            throw new Error('missing DocSearch apiKey');
        }
        if (!cfg.indexName) {
            throw new Error('missing DocSearch indexName');
        }

        docsearch({
            appId: cfg.appId,
            apiKey: cfg.apiKey,
            indexName: cfg.indexName,
            debug: cfg.debug ?? false,
            container: '#book-doc-search-input',
        });
    }
    gitbook.events.bind("start", function(e, config) {
        pluginsConfig = config;
        initDocSearch();
    });
    gitbook.events.bind("page.change", function() {
        initDocSearch();
    });
});
