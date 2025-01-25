// utils/format.ts
export const formatDuration = (duration: string) => {
    const time = duration.replace('m', ' min').replace('s', '');
    return time.endsWith(' min') ? time : `${time} sec`;
  };

  export const formatDistance = (distance: number) => {
    return distance.toFixed(2);
  };
