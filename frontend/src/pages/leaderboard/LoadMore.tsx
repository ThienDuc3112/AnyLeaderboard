import React, { useEffect, useRef } from "react";

type PropType = {
  fn: () => void;
  hasMore: boolean;
};

const LoadMore: React.FC<PropType> = ({ fn, hasMore }) => {
  const ref = useRef<HTMLDivElement>(null);
  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          if (hasMore) fn();
          else observer.disconnect();
        }
      },
      {
        threshold: 0.5,
      },
    );
    if (ref.current) observer.observe(ref.current);
    return () => observer.disconnect();
  }, []);

  if (!hasMore)
    return (
      <div className="w-full text-slate-500 text-center my-4">
        There's no more leaderboards
      </div>
    );
  return (
    <div ref={ref} className="w-full text-slate-500 text-center my-4">
      Loading more...
    </div>
  );
};

export default LoadMore;
