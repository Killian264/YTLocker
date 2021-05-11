import React from "react";
import { Playlist } from "../shared/types";
import { Link } from "./Link";

export interface PlaylistItemProps {
	className?: string;
	playlist: Playlist;
}

export const PlaylistListItem: React.FC<PlaylistItemProps> = ({
	playlist,
	className,
}) => {
	return (
		<div
			className={`${className} font-semibold text-xl hover:bg-secondary-hover rounded-md flex justify-between`}
		>
			<div className="flex">
				<img
					src={playlist.thumbnail}
					alt="Logo"
					width="140"
					height="120"
					className="rounded-lg"
				/>
				<div className="pl-3 flex flex-col">
					<span>{playlist.title}</span>
					<Link
						className="text-accent text-lg ml-0.5"
						href={playlist.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-3 text-3xl my-auto select-none">{">"}</div>
		</div>
	);
};
