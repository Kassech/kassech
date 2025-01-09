/* eslint-disable @typescript-eslint/no-explicit-any */
import { AxiosError } from "axios";
import {
  useQuery,
  useMutation,
  UseMutationOptions,
  UseMutationResult,
} from "react-query";

// Utility function to handle API response standardization
const handleApiResponse = (response: any) => {
  if (response.status === "success") {
    return response.data; // Return the data part of the response if it's a success
  } else {
    // Extract the error message from response and throw it as a string
    throw new Error(
      response.message || response.data?.error || "An unknown error occurred"
    );
  }
};

// Custom Query Hook
export const useCustomQuery = (
  key: string,
  fetchFn: () => Promise<any>,
  options = {}
) =>
  useQuery(
    key,
    async () => {
      const response = await fetchFn(); // Fetch the response using the provided function
      return handleApiResponse(response); // Standardize the response format
    },
    options
  );
  export const useCustomMutation = <
  TRequest, // Request data type (input)
  TResponse, // Response data type (success)
  TError = any // Error type, default to any
>(
  mutateFn: (data: TRequest) => Promise<TResponse>, // The function that performs the mutation
  options: UseMutationOptions<TResponse, TError, TRequest> = {} // Error type is dynamic
): UseMutationResult<TResponse, TError, TRequest> =>
  useMutation<TResponse, TError, TRequest>(mutateFn, {
      ...options,
      onError: (error: TError) => {
          const axiosError = error as AxiosError;
          let errorMessage = "An unknown error occurred";

          if (axiosError.response) {
              const status = axiosError.response.status;
              const data = axiosError.response.data as any;

              if (status === 400 && data.errors) {
                  errorMessage = Object.values(data.errors).join(", ");
              } else if (status === 401) {
                  errorMessage = data.error || "Invalid credentials";
              } else if (status === 500) {
                  errorMessage = data.error || "Internal server error";
              } else {
                  errorMessage = data.error || "Unexpected error";
              }
          } else if (axiosError.message) {
              errorMessage = axiosError.message;
          }

          console.error("🚀 ~ Error:", errorMessage);
          console.error("🚀 ~ Error:", errorMessage);
      },
      onSuccess: (data) => {
          console.log("Mutation succeeded with data:", data);
      },
  });
