"use client";
import React from "react";
import { useDropzone } from "react-dropzone";
import { Input } from "@/components/ui/input";
import { ImagePlus } from "lucide-react";
import { Avatar } from "./ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";

interface ImageUploaderProps {
    onImageUpload: (file: File) => void;
    maxFileSize?: number; // Default: 1MB
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    acceptedFormats?: Record<string, any>; // Default: PNG, JPG, JPEG
}

const ImageUploader: React.FC<ImageUploaderProps> = ({
    onImageUpload,
    maxFileSize = 1000000,
    acceptedFormats = { "image/png": [], "image/jpg": [], "image/jpeg": [] },
}) => {
    const [preview, setPreview] = React.useState<string | ArrayBuffer | null>("");

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
        [onImageUpload],
    );

    const { getRootProps, getInputProps, isDragActive, fileRejections } =
        useDropzone({
            onDrop,
            maxFiles: 1,
            maxSize: maxFileSize,
            accept: acceptedFormats,
        });

    return (
        <div className="flex flex-col items-center justify-center">
            <Avatar
                {...getRootProps()}
                className="h-40 w-40 flex items-center justify-center border-2 border-dashed border-gray-300 rounded-full"
            >
                {preview && (
                    <AvatarImage
                        src={preview as string}
                        alt="Uploaded image"
                        className="h-full w-full object-cover rounded-full"
                    />
                )}
                <ImagePlus className={`h-10 w-10 ${preview ? "hidden" : "block"}`} />
                <Input {...getInputProps()} type="file" className="hidden" />
            </Avatar>
            {isDragActive ? (
                <p className="mt-2 text-center">Drop the image!</p>
            ) : (
                <p className="mt-2 text-center">Click here or drag an image to upload it</p>
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
