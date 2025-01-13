import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
export function base64ToFile(base64String) {
    const arr = base64String.split(',');
    const mime = arr[0].match(/:(.*?);/)[1];  // Extract MIME type
    const bstr = atob(arr[1]);  // Decode base64 to binary string
    let n = bstr.length;
    const u8arr = new Uint8Array(n);
    while (n--) {
      u8arr[n] = bstr.charCodeAt(n);  // Convert binary to Uint8Array
    }
    return new File([u8arr], 'profile.jpg', { type: mime });  // Create File object
  };
