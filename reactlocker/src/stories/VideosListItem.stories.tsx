import { Story } from "@storybook/react";
import { Card } from "../components/Card";
import { VideoListItem } from "../components/VideosListItem";
import { PlusButton } from "../components/PlusButton";
import { Video } from "../shared/types";

export default {
	title: "VideoListItem",
	component: VideoListItem,
};

const videos: Video[] = [
	{
		id: 932423423,
		youtube: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnail:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		url:
			"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
		created: new Date(),
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	return (
		<Card>
			<div></div>
			<div>
				<VideoListItem video={videos[0]}></VideoListItem>
				<VideoListItem
					video={videos[0]}
					className="mt-3"
				></VideoListItem>
				<VideoListItem
					video={videos[0]}
					className="mt-3"
				></VideoListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
