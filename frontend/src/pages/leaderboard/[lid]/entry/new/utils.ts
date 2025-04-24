export const parseDurationToMs = (str: string): number => {
  if (!str || typeof str !== "string") return 0;

  const [timePart, msPart] = str.trim().split(".");
  const timeParts = timePart.split(":").reverse();

  if (timeParts.some((p) => isNaN(Number(p)))) return 0;
  if (msPart && isNaN(Number(msPart))) return 0;

  const seconds = parseInt(timeParts[0] || "0", 10);
  const minutes = parseInt(timeParts[1] || "0", 10);
  const hours = parseInt(timeParts[2] || "0", 10);
  const milliseconds = parseInt(msPart || "0", 10);

  return hours * 3600000 + minutes * 60000 + seconds * 1000 + milliseconds;
};

export const formatMsToDuration = (ms: number): string => {
  const hours = Math.floor(ms / 3600000);
  const minutes = Math.floor((ms % 3600000) / 60000);
  const seconds = Math.floor((ms % 60000) / 1000);
  const milliseconds = ms % 1000;

  const parts = [
    hours.toString().padStart(2, "0"),
    minutes.toString().padStart(2, "0"),
    seconds.toString().padStart(2, "0"),
  ];

  return `${parts.join(":")}.${milliseconds.toString().padStart(3, "0")}`;
};

export const parseTimestampToMs = (input: string): number => {
  if (!input || typeof input !== "string") return 0;

  // Normalize to ISO format: replace space with "T"
  const normalized = input.trim().replace(" ", "T");

  // Try parse
  const date = new Date(normalized);

  return isNaN(date.valueOf()) ? 0 : date.valueOf();
};

export const formatMsToTimestamp = (ms: number): string => {
  const date = new Date(ms);
  const pad = (n: number) => n.toString().padStart(2, "0");

  const yyyy = date.getFullYear();
  const mm = pad(date.getMonth() + 1);
  const dd = pad(date.getDate());
  const hh = pad(date.getHours());
  const min = pad(date.getMinutes());
  const ss = pad(date.getSeconds());

  return `${yyyy}-${mm}-${dd} ${hh}:${min}:${ss}`;
};
