import React from "react";
import { BuildChannelUrl } from "../shared/urls";
import { ChannelListItem } from "../components/ChannelListItem";
import { useChannel } from "../hooks/api/useChannel";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { Color } from "../shared/types";
import { LoadingListItem } from "../components/LoadingListItem";

export interface ChannelListItemControllerProps {
	className?: string;
	channelId: number;
	mode: "normal" | "delete";
	remove: (channelId: string) => void;
}

export const ChannelListItemController: React.FC<ChannelListItemControllerProps> = ({
	className,
	channelId,
	mode = "normal",
	remove,
}) => {
	const [isLoadingChannel, channel] = useChannel(channelId);
	const [isLoadingPlaylist, playlists] = usePlaylistList();

	if (isLoadingChannel || channel == null) {
		return <LoadingListItem></LoadingListItem>;
	}

	if (isLoadingPlaylist || playlists == null) {
		return <LoadingListItem></LoadingListItem>;
	}

	let colors: Color[] = [];
	playlists.forEach((playlist) => {
		if (playlist.channels.includes(channelId)) {
			colors.push(playlist.color);
		}
	});

	return (
		<ChannelListItem
			className={className}
			channel={channel}
			url={BuildChannelUrl(channel.youtubeId)}
			colors={colors}
			mode={mode}
			remove={remove}
		></ChannelListItem>
	);
};
