---
layout: default
title: React useContext with useReducer
nav_exclude: true
parent: Browser JS
grand_parent: Programming Languages
permalink: /docs/languages/node-js/framework-react
---

__[back]({% link docs/languages/js/framework.md %})__

## React useContext with useReducer


<details markdown="block">
  <summary>
    code snippet
  </summary>

A sample from a DEV Community post [^1].

```bash
# init
npx create-react-app@5.0.1 react-context-with-usereducer
cd react-context-with-usereducer
```


<details markdown="block">
  <summary>
    <sup>app-context.js</sup>
  </summary>
```jsx
import React, { useMemo, useReducer, createContext, useContext } from 'react';
import { initialState, contextReducer } from './app-context-reducer';
import contextActions from './app-context-actions';

const AppContext = createContext();

function AppContextProvider({ children }) {
  const [state, dispatch] = useReducer(contextReducer, initialState);
  const value = useMemo(() => [state, dispatch], [state]);

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
}

function useAppContext() {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error('useAppContext must be used within a AppContextProvider');
  }
  const [state, dispatch] = context;
  const appContextAction = contextActions(dispatch);
  // const appContextSelector = contextSelectors(state);
  return { state, appContextAction };
}

export { AppContextProvider, useAppContext };
```
------
<!-- app-context.js -->
</details>

<details markdown="block">
  <summary>
    <sup>app-context-types.js</sup>
  </summary>
```js
export const OPEN_NAV_MENU = 'OPEN_NAV_MENU';
export const CLOSE_NAV_MENU = 'CLOSE_NAV_MENU';
export const COLLAPSE_NAV_MENU = 'COLLAPSE_NAV_MENU';
```
------
<!-- app-context-types.js -->
</details>

<details markdown="block">
  <summary>
    <sup>app-context-reducer.js</sup>
  </summary>
```js
import * as actionTypes from './app-context-types';

// Define the initial state for the context
export const initialState = {
  isNavMenuClose: false,
};

// Define the reducer function for the context
export function contextReducer(state, action) {
  switch (action.type) {
    // Handle the OPEN_NAV_MENU action
    case actionTypes.OPEN_NAV_MENU:
      return {
        ...state,
        isNavMenuClose: false,
      };
    // Handle the CLOSE_NAV_MENU action
    case actionTypes.CLOSE_NAV_MENU:
      return {
        ...state,
        isNavMenuClose: true,
      };
    // Handle the COLLAPSE_NAV_MENU action
    case actionTypes.COLLAPSE_NAV_MENU:
      return {
        ...state,
        isNavMenuClose: !state.isNavMenuClose,
      };
    // Throw an error for any unhandled action types
    default: {
      throw new Error(`Unhandled action type: ${action.type}`);
    }
  }
}
```
------
<!-- app-context-reducer.js -->
</details>

<details markdown="block">
  <summary>
    <sup>app-context-actions.js</sup>
  </summary>
```js
import * as actionTypes from './app-context-types';

// Define a function that returns an object with context actions
const contextActions = (dispatch) => {
  return {
    navMenu: {
      // Action for opening the navigation menu
      open: () => {
        dispatch({ type: actionTypes.OPEN_NAV_MENU });
      },
      // Action for closing the navigation menu
      close: () => {
        dispatch({ type: actionTypes.CLOSE_NAV_MENU });
      },
      // Action for toggling (collapsing/expanding) the navigation menu
      collapse: () => {
        dispatch({ type: actionTypes.COLLAPSE_NAV_MENU });
      },
    },
  };
};

export default contextActions;
```
------
<!-- app-context-actions.js -->
</details>

<details markdown="block">
  <summary>
    <sup>App.js</sup>
  </summary>
```jsx
import logo from './logo.svg';
import './App.css';
import { AppContextProvider } from './app-context';
import Sample from './Sample';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <AppContextProvider>
          <Sample />
        </AppContextProvider>
        <br/>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
```
------
<!-- App.js -->
</details>

<details markdown="block">
  <summary>
    <sup>App.css</sup>
  </summary>
```css
.App {
  text-align: center;
}

.App-logo {
  height: 40vmin;
  pointer-events: none;
}

@media (prefers-reduced-motion: no-preference) {
  .App-logo {
    animation: App-logo-spin infinite 20s linear;
  }
}

.App-header {
  background-color: #282c34;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 2vmin);
  color: white;
}

.App-link {
  color: #61dafb;
}

@keyframes App-logo-spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
```
------
<!-- App.css -->
</details>

<details markdown="block">
  <summary>
    <sup>index.js</sup>
  </summary>
```jsx
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
```
------
<!-- index.js -->
</details>

<details markdown="block">
  <summary>
    <sup>index.css</sup>
  </summary>
```css
body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New',
    monospace;
}
```
------
<!-- index.css -->
</details>

<details markdown="block">
  <summary>
    <sup>Sample.js</sup>
  </summary>
```jsx
import { useAppContext } from './app-context';

function Sample() {
  const { state: stateApp, appContextAction } = useAppContext();
  const { isNavMenuClose } = stateApp;
  const { navMenu } = appContextAction;

  const onCollapse = () => navMenu.collapse();
  const onOpen = () => navMenu.open();
  const onClose = () => navMenu.close();

  return (
    <>
      <span>{'my state:'}</span>
      <span>{`${JSON.stringify(stateApp)}`}</span>
      <button type="button" onClick={onOpen}>Open</button>
      <button type="button" onClick={onClose}>Close</button>
      <button type="button" onClick={onCollapse}>Collapse</button>
    </>
  );
}

export default Sample;
```
------
<!-- Sample.js -->
</details>

After creating all files, the repo should look like below.

```
 |-src
 | |-app-context.js
 | |-app-context-types.js
 | |-app-context-reducer.js
 | |-app-context-actions.js
 | |-App.js
 | |-App.css
 | |-index.js
 | |-index.css
 | |-Sample.js
```

```bash
# begin by typing:
cd react-context-with-usereducer
npm start

# starts the development server.
npm start

# bundles the app into static files for production.
npm run build

# starts the test runner.
npm test

# removes this tool and copies build dependencies, configuration files
# and scripts into the app directory. If you do this, you can’t go back!
npm run eject
```

```bash
# GET request
curl localhost:3000
```

</details>


----

[^1]: [Mastering Advanced Complex React useContext with useReducer (Redux-like Style)](https://dev.to/idurar/mastering-advanced-complex-react-usecontext-with-usereducer-redux-style-2jl0)
