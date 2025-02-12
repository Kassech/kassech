import { cn } from '@/lib/utils';
import { Loader } from "lucide-react";

const LoadingSpinner = ({ className }: { className?: string }) => {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <Loader
        className={cn('h-16 w-16 text-primary/60 animate-spin', className)}
      />
    </div>
  );
};

export default LoadingSpinner;
