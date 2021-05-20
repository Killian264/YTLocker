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
