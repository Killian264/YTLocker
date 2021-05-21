import React from "react";
import { BuildChannelUrl } from "../shared/urls";
import { ChannelListItem } from "../components/ChannelListItem";
import { useChannel } from "../shared/api/useChannel";
import { usePlaylists } from "../shared/api/usePlaylists";
import { Color } from "../shared/types";

export interface ChannelListItemControllerProps {
	className?: string;
	channelId: number;
}

export const ChannelListItemController: React.FC<ChannelListItemControllerProps> = ({
	className,
	channelId,
}) => {
	const [loadingC, channel] = useChannel(channelId);
	const [loadingP, playlists] = usePlaylists();

	if (loadingC || channel == null) {
		return <div>Loading...</div>;
	}

	if (loadingP || playlists == null) {
		return <div>Loading...</div>;
	}

	let colors: Color[] = [];
	playlists.forEach((playlist) => {
		if (playlist.channels.includes(channelId)) {
			colors.push(playlist.color);
		}
	});

	return (
		<ChannelListItem
			channel={channel}
			url={BuildChannelUrl(channel.youtubeId)}
			playlistColors={colors}
		></ChannelListItem>
	);
};
