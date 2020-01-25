import React, { PureComponent } from 'react';

export const setLocalStorageItem = (key: string, value: any, notifyListener = false) => {
  localStorage.setItem(key, JSON.stringify(value));
  if (notifyListener) {
    window.dispatchEvent(
      new StorageEvent('storage', {
        key,
        newValue: JSON.stringify(value),
      })
    );
  }
};

export const getLocalStorageItem = <T extends {}>(key: string, fallback: string): T => {
  return JSON.parse(localStorage.getItem(key) || fallback);
};

class LocalStorageListener extends PureComponent<
  { children: (storage: Record<string, any>) => JSX.Element; state: Record<string, string> },
  any
> {
  state = Object.entries(this.props.state as Record<string, string>).reduce((acc, [key, fallback]) => {
    return {
      ...acc,
      [key]: getLocalStorageItem(key, fallback),
    };
  }, {});
  componentDidMount() {
    window.addEventListener('storage', e => {
      if (e.key) {
        this.setState({
          ...this.state,
          [e.key]: JSON.parse(e.newValue!),
        });
      }
    });
  }
  render() {
    return <>{this.props.children(this.state)}</>;
  }
}

export default LocalStorageListener;
