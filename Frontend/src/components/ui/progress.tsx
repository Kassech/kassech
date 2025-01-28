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
      className={cn('relative w-24 h-24', className)} // Fixed size for the circular progress bar
      {...props}
    >
      <svg
        className="absolute top-0 left-0 transform -translate-x-1/2 -translate-y-1/2"
        width="100%"
        height="100%"
        viewBox="0 0 120 120"
      >
        <circle
          cx="60"
          cy="60"
          r={radius}
          stroke="currentColor"
          strokeWidth={strokeWidth}
          fill="none"
          className="text-secondary" // Background circle
        />
        <circle
          cx="60"
          cy="60"
          r={radius}
          stroke="currentColor"
          strokeWidth={strokeWidth}
          fill="none"
          className="text-primary" // Foreground circle (progress)
          strokeDasharray={circumference}
          strokeDashoffset={strokeDashoffset}
          style={{ transition: 'stroke-dashoffset 0.3s ease' }}
        />
      </svg>
      <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-xl font-bold text-black">
        {value}%
      </div>
    </ProgressPrimitive.Root>
  );
});

Progress.displayName = ProgressPrimitive.Root.displayName;

export { Progress };
