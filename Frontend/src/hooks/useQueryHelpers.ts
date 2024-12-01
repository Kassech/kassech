/* eslint-disable @typescript-eslint/no-explicit-any */
import { useQuery, useMutation } from 'react-query';

export const useCustomQuery = (key: string, fetchFn: () => Promise<any>, options = {}) =>
  useQuery(key, fetchFn, options);

export const useCustomMutation = (mutateFn: (data: any) => Promise<any>, options = {}) =>
  useMutation(mutateFn, options);
