import { useState, useEffect } from 'react';

export type APIResponse<T> = { status?: string; data?: T };

export interface FetchState<T> {
  data?: T;
  status?: string;
  error?: Error;
  isLoading: boolean;
}

const initialResponseState: APIResponse<any> = { status: undefined, data: undefined };

export const useFetch = <T extends {}>(url: string, options?: RequestInit): FetchState<T> => {
  const [response, setResponse] = useState<APIResponse<T>>(initialResponseState);
  const [error, setError] = useState<Error>();
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const res = await fetch(url, options);
        if (!res.ok) {
          throw new Error(res.statusText);
        }
        const json = (await res.json()) as APIResponse<T>;
        setResponse(json || initialResponseState);
        setIsLoading(false);
      } catch (error) {
        setError(error);
      }
    };
    fetchData();
  }, [url, options]);
  return { ...response, error, isLoading };
};
