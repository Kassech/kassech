'use client';
import React from 'react';
import { useDropzone } from 'react-dropzone';
import { Input } from '@/components/ui/input';
import { ImagePlus } from 'lucide-react';
import { Avatar } from './ui/avatar';
import { AvatarImage } from '@radix-ui/react-avatar';
import { apiEndpoint } from '@/config/config';

interface ImageUploaderProps {
  onImageUpload: (file: File) => void;
  maxFileSize?: number; // Default: 1MB
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  acceptedFormats?: Record<string, any>; // Default: PNG, JPG, JPEG
  initialPreview?: File | string | null; // Initial preview image (File or base64 string)
  className?: string;
}

const ImageUploader: React.FC<ImageUploaderProps> = ({
  onImageUpload,
  maxFileSize = 1000000,
  acceptedFormats = { 'image/png': [], 'image/jpg': [], 'image/jpeg': [] },
  initialPreview = null,
  className = 'rounded-full',
}) => {
  const [preview, setPreview] = React.useState<string | ArrayBuffer | null>(
    null
  );

  React.useEffect(() => {
    if (initialPreview) {
      if (typeof initialPreview === 'string') {
        setPreview(initialPreview); // If it's already a base64 string
      } else if (initialPreview instanceof File) {
        const reader = new FileReader();
        reader.onload = () => setPreview(reader.result);
        reader.readAsDataURL(initialPreview);
      }
    }
  }, [initialPreview]);

  const onDrop = React.useCallback(
    (acceptedFiles: File[]) => {
      const reader = new FileReader();
      try {
        reader.onload = () => setPreview(reader.result);
        reader.readAsDataURL(acceptedFiles[0]);
        onImageUpload(acceptedFiles[0]);
      } catch {
        setPreview(null);
      }
    },
    [onImageUpload]
  );

  const { getRootProps, getInputProps, isDragActive, fileRejections } =
    useDropzone({
      onDrop,
      maxFiles: 1,
      maxSize: maxFileSize,
      accept: acceptedFormats,
    });
  console.log('🚀 ~ preview:', preview);
  return (
    <div className="flex flex-col items-center justify-center">
      <Avatar
        {...getRootProps()}
        className={`h-40 w-40 flex items-center justify-center border-2 border-dashed border-gray-300 ${className}`}
      >
        {preview && (
          <AvatarImage
            // Check if the preview string starts with "http"
            src={
              preview.startsWith('http') ? preview : `${apiEndpoint}/${preview}`
            }
            alt="Uploaded image"
            className={`h-full w-full object-cover ${className}`}
          />
        )}
        <ImagePlus className={`h-10 w-10 ${preview ? 'hidden' : 'block'}`} />
        <Input {...getInputProps()} type="file" className="hidden" />
      </Avatar>

      {isDragActive ? (
        <p className="mt-2 text-center">Drop the image!</p>
      ) : (
        <p className="mt-2 text-center">Click or drag an image</p>
      )}
      {fileRejections.length !== 0 && (
        <p className="mt-2 text-center text-destructive">
          Image must be less than {maxFileSize / 1000000}MB and of accepted type
        </p>
      )}
    </div>
  );
};

export default ImageUploader;
