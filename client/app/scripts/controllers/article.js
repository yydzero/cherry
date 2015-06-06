'use strict';

angular.module('clientApp').controller('ArticleCtrl', ['$scope', '$http', '$routeParams', '$log', '$sce', 'config', function ($scope, $http, $routeParams, $log, $sce, config) {
    $scope.article = {};
    $scope.fullContent = '';
    $scope.message = '';
    $scope.deleting = false;
    $scope.config = config;

    function stripScripts(s) {
        var div = document.createElement('div');
        div.innerHTML = s;
        var scripts = div.getElementsByTagName('script');
        var i = scripts.length;
        while (i--) {
            scripts[i].parentNode.removeChild(scripts[i]);
        }
        return div.innerHTML;
    }

    $scope.getArticlesById = function(id) {
        $http.get(config.url + '/article/' + id)
            .success(function (resp) {
                if (resp.Status === 'ok') {
                    $scope.article = resp.Resource;
                    var c = stripScripts($scope.article.FullContent);
                    $scope.fullContent = $sce.trustAsHtml(c);
                    $scope.message = '';
                } else {
                    $scope.message = resp.Message;
                }
            }).error(function (error) {
                $scope.message = '不能获得订阅的公众号列表.';
            }
        );
    };


    $scope.getArticlesById($routeParams.id);

}]);