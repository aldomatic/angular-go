(function() {
  'use strict';

  angular
    .module('angularGo')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($timeout, $log) {
    var vm = this;
    vm.title = "Angular and Go!!";
    activate();


    function activate() {
      $timeout(function() {
        $log.debug("active()");
      }, 3000);
    }
  }
})();
