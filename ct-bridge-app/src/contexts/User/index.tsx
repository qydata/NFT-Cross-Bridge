// contexts/User/index.jsx
import React from 'react';
import { initialState, reducer } from './reducer';

export const UserContext = React.createContext<any>({
  state: initialState,
  dispatch: () => null,
  value: []
});
export const UserProvider = ({ children }: any) => {
  const [state, dispatch] = React.useReducer(reducer, initialState);

  return (
    <UserContext.Provider value={[state, dispatch]}>
      {children}
    </UserContext.Provider>
  );
};
