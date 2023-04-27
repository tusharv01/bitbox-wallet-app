/**
 * Copyright 2020 Shift Crypto AG
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* global jest */

// When a module imports 'react-i18next', it gets replaced with this mock
// by Jest in tests.
// It allows for mounting components with stubs covering lifecycle methods.


module.exports = reactI18next;
import React from 'react';
import { Suspense } from 'react';
import { useTranslation, I18nextProvider } from 'react-i18next';
import i18n from 'i18next';

// Remove any calls to i18n.addResourceBundle or similar methods here

function App() {
  const { t } = useTranslation('common', { useSuspense: false });

  // Load additional language resources on demand as needed
  function loadLanguageResources(lang) {
    return i18n.loadNamespaces(lang, ['common']);
  }

  return (
    <I18nextProvider i18n={i18n}>
      <Suspense fallback={<div>Loading...</div>}>
        {/* Render your application here */}
      </Suspense>
    </I18nextProvider>
  );
}

export default App;

