import * as React from 'react';
import * as ProgressPrimitive from '@radix-ui/react-progress';
import { cn } from '@/lib/utils';

const Progress = React.forwardRef<
  React.ElementRef<typeof ProgressPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof ProgressPrimitive.Root>
>(({ className, value, ...props }, ref) => {
  const radius = 50; // Radius of the circle
  const strokeWidth = 8; // Stroke width for the circle
  const circumference = 2 * Math.PI * radius; // Circumference of the circle
  const strokeDashoffset = circumference - (value || 0) * (circumference / 100);

  return (
    <ProgressPrimitive.Root
      ref={ref}
      className={cn(
        'relative flex items-center justify-center w-32 h-32',
        className
      )}
      {...props}
    >
      <svg width="100%" height="100%" viewBox="0 0 120 120" className="block">
        {/* Background Circle */}
        <circle
          cx="60"
          cy="60"
          r={radius}
          stroke="currentColor"
          strokeWidth={strokeWidth}
          fill="none"
          className="text-gray-300"
        />
        {/* Progress Circle */}
        <circle
          cx="60"
          cy="60"
          r={radius}
          stroke="currentColor"
          strokeWidth={strokeWidth}
          fill="none"
          className="text-primary"
          strokeDasharray={circumference}
          strokeDashoffset={strokeDashoffset}
          strokeLinecap="round"
          style={{ transition: 'stroke-dashoffset 0.5s ease' }}
        />
      </svg>
      {/* Centered Text */}
      <div className="absolute text-xl font-bold text-black">{value}%</div>
    </ProgressPrimitive.Root>
  );
});

Progress.displayName = 'Progress';

export { Progress };
