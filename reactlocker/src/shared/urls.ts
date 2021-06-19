export const BuildVideoUrl = (videoId: string): string => {
	return `https://www.youtube.com/watch?v=${videoId}`;
};

export const BuildVideoPlaylistUrl = (videoId: string, playlistId: string): string => {
	return `https://www.youtube.com/watch?v=${videoId}&list=${playlistId}`;
};

export const BuildChannelUrl = (channelId: string): string => {
	return `https://www.youtube.com/channel/${channelId}`;
};

export const BuildPlaylistUrl = (playlistId: string): string => {
	return `https://www.youtube.com/playlist?list=${playlistId}`;
};

// url types:
// https://www.youtube.com/channel/UCCfqyGl3nq_V0bo64CjZh8g
// https://www.youtube.com/c/ContinuousDelivery
// returns kind: ["username" | "id", "parsed id"]
export const ParseYTChannelUrl = (url: string): [string, string] => {
	let parts = url.trim().replace("https://www.youtube.com/", "").split("/");

	let id = parts[1];
	let kind = parts[0] === "channel" ? "id" : "username";

	return [kind, id];
};

export const IsValidYTChannelUrl = (url: string): boolean => {
	return url.includes("https://www.youtube.com/channel/") || url.includes("https://www.youtube.com/c/");
};
