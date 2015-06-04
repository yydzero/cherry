'use strict';


// TODO: 如何在多个 controller 之间共享数据，其中某个数据变化后，多个controller同时更新view？
// 譬如添加或者删除公众号， 文章列表也同时更新。
angular.module('clientApp').factory('gzhService', ['$http', '$log', '$q', 'config', function ($http, $log, $q, config) {

    return {
        isCrawlLoading: false,
        crawlStats: [],
        crawlMessage: '',

        // 从数据库中读取每个公众号的抓取状态，包括文章数、日期等。
        getGzhCrawlStatus: function () {
            var deferred = $q.defer();

            $http.get(config.url + '/stats').success(function (resp) {
                if (resp.Status === 'ok') {
                    self.crawlStats = resp.Resource;
                    self.crawlMessage = '';
                    $log.info(self.crawlStats);

                    deferred.resolve(self.crawlStats);
                } else {
                    self.crawlMessage = resp.Message
                }
            }).error(function (error) {
                deferred.reject({message: "Really bad"});
            });

            return deferred.promise;
        },

        crawlGzh: function(gzh) {
            var deferred = $q.defer();
            var gzhStats = {};

            $http.get(config.url + '/crawl/' + gzh.OpenId).success(function (resp) {
                if (resp.Status === 'ok') {
                    gzhStats = resp.Resource;

                    deferred.resolve(gzhStats);
                } else {
                    console.log(resp.Message);
                }
            }).error(function (error) {
                console.log(error);
                deferred.reject({message: "Really bad"});
            });

            return deferred.promise;
        }

        //getGzhCrawlStatus: function() {
        //    var self = this;
        //    $log.info("getting crawl status");
        //    self.isCrawlLoading = true;
        //    $http.get(config.url + '/stats')
        //        .success(function (resp) {
        //            $log.info(resp);
        //            if (resp.Status === 'ok') {
        //                self.crawlStats = resp.Resource;
        //                self.crawlMessage = '';
        //                $log.info(self.crawlStats);
        //            } else {
        //                self.crawlMessage = resp.Message
        //            }
        //            self.isCrawlLoading = false;
        //        }).error(function (error) {
        //            self.crawlMessage = '不能获得订阅的公众号爬虫状态信息.';
        //            self.isCrawlLoading = false;
        //        });
        //}
    };
}]);