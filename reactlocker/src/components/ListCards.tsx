import { Channel, Playlist, Video } from "../shared/types";
import { BuildPlaylistUrl, BuildVideoUrl } from "../shared/urls";
import { Card } from "./Card";
import { ChannelListItem } from "./ChannelListItem";
import { PlaylistListItem } from "./PlaylistListItem";
import { PlusButton } from "./PlusButton";
import { VideoListItem } from "./VideosListItem";

interface PlaylistListProps {
	className?: string;
	playlists: Playlist[];
	limit: number;
}

export const PlaylistList: React.FC<PlaylistListProps> = ({ className, playlists, limit }) => {
	let list = playlists.map((playlist, index) => {
		if (index >= limit) {
			return "";
		}

		return (
			<PlaylistListItem
				url={BuildPlaylistUrl(playlist.youtubeId)}
				key={index}
				playlist={playlist}
			></PlaylistListItem>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<PlusButton color="secondary" disabled={false}></PlusButton>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface VideoListProps {
	className?: string;
	videos: Video[];
	limit: number;
}

export const VideoList: React.FC<VideoListProps> = ({ className, videos, limit }) => {
	let list = videos.map((video, index) => {
		if (index >= limit) {
			return "";
		}

		return <VideoListItem url={BuildVideoUrl(video.youtubeId)} key={index} video={video}></VideoListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface ChannelsListProp {
	className?: string;
	channels: number[];
	limit: number;
}

export const ChannelsList: React.FC<ChannelsListProp> = ({ className, channels, limit }) => {
	// let list = channels.map((channel, index) => {
	// 	if (index >= limit) {
	// 		return "";
	// 	}

	// 	return <ChannelListItem key={index} channel={channel}></ChannelListItem>;
	// });

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
			</div>
			{/* <div className="grid gap-2">{list}</div> */}
			<div></div>
		</Card>
	);
};
