(function() {
  'use strict';

  angular
    .module('angularGo')
    .run(runBlock);

  /** @ngInject */
  function runBlock($log) {

    $log.debug('runBlock end');
  }

})();
