import React, { useState } from "react";
import { Card } from "../components/Card";
import { ChannelListItemController } from "./ChannelListItemController";
import { Playlist } from "../shared/types";
import { Cog, SvgBox } from "../components/Svg";

export interface PlaylistChannelListControllerProps {
	className?: string;
	limit: number;
	playlist: Playlist;
}

export const PlaylistChannelListController: React.FC<PlaylistChannelListControllerProps> = ({
	className,
	limit,
	playlist,
}) => {
	const [mode, setMode] = useState<"normal" | "delete">("normal");

	const list = playlist.channels.map((id, index) => {
		if (index >= limit) {
			return "";
		}
		return (
			<ChannelListItemController
				mode={mode}
				remove={(id: number) => {
					console.log(id);
				}}
				key={index}
				channelId={id}
			></ChannelListItemController>
		);
	});

	const swap = () => {
		setMode(mode === "normal" ? "delete" : "normal");
	};

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
				<SvgBox className="text-primary-200 p-0.5 border-primary-200" onClick={swap}>
					<Cog size={26}></Cog>
				</SvgBox>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
