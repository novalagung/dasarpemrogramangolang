require([
    "gitbook",
    "jQuery"
], function(gitbook, $) {

    function removeDisqus() {
        setTimeout(function () {
            var contents = [
                "A. Pemrograman Golang Dasar",
                "B. Pemrograman Web Golang Dasar",
                "C. Pemrograman Web Golang Lanjut"
            ];
            if (contents.indexOf(gitbook.state.page.title) > -1) {
                $('#disqus_thread').remove();
            }

            contents.forEach(function (each) {
                $('a:contains("' + each + '")').addClass('no-link')
            });
        }, 500)
    }

    gitbook.events.bind("start", removeDisqus);
    gitbook.events.bind("page.change", removeDisqus);
});
