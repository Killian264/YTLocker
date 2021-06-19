import React from "react";
import { BuildChannelUrl } from "../shared/urls";
import { ChannelListItem } from "../components/ChannelListItem";
import { useChannel } from "../hooks/api/useChannel";
import { usePlaylists } from "../hooks/api/usePlaylists";
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
	const [loadingC, channel] = useChannel(channelId);
	const [loadingP, playlists] = usePlaylists();

	if (loadingC || channel == null) {
		return <LoadingListItem></LoadingListItem>;
	}

	if (loadingP || playlists == null) {
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
