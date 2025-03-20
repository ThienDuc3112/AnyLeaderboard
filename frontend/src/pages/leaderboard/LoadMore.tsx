import React, { useEffect } from 'react'

type PropType = {
  fn: () => void
  hasMore: boolean
}

const LoadMore: React.FC<PropType> = ({ fn, hasMore }) => {
  useEffect(() => {
    if (hasMore) fn()
    return () => { }
  }, [])

  if (!hasMore) return;
  return (
    <div>LoadMore</div>
  )
}

export default LoadMore
