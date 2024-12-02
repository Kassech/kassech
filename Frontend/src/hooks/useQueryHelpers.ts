/* eslint-disable @typescript-eslint/no-explicit-any */
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
  string // Error type
>(
  mutateFn: (data: TRequest) => Promise<TResponse>, // The function that performs the mutation
  options: UseMutationOptions<TResponse, string, TRequest> = {} // Error type is explicitly string
): UseMutationResult<TResponse, string, TRequest> =>
  useMutation<TResponse, string, TRequest>(mutateFn, {
    ...options,
    onError: (error: unknown) => {
      const errorMessage =
        error?.response?.data?.error ||
        error.message ||
        "An unknown error occurred";
      // Pass the error as a string to React Query
      console.log("🚀 ~ errorMessage:", errorMessage);

      return errorMessage;
    },
    onSuccess: (data) => {
      console.log("Mutation succeeded with data:", data);
    },
  });
