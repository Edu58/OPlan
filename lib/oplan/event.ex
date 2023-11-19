defmodule Oplan.Event do
  @moduledoc """
  The Event context.
  """

  import Ecto.Query, warn: false
  alias Oplan.Repo

  alias Oplan.Event.EventRating

  @doc """
  Returns the list of event_rating.

  ## Examples

      iex> list_event_rating()
      [%EventRating{}, ...]

  """
  def list_event_rating do
    Repo.all(EventRating)
  end

  @doc """
  Gets a single event_rating.

  Raises `Ecto.NoResultsError` if the Event rating does not exist.

  ## Examples

      iex> get_event_rating!(123)
      %EventRating{}

      iex> get_event_rating!(456)
      ** (Ecto.NoResultsError)

  """
  def get_event_rating!(id), do: Repo.get!(EventRating, id)

  @doc """
  Creates a event_rating.

  ## Examples

      iex> create_event_rating(%{field: value})
      {:ok, %EventRating{}}

      iex> create_event_rating(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_event_rating(attrs \\ %{}) do
    %EventRating{}
    |> EventRating.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a event_rating.

  ## Examples

      iex> update_event_rating(event_rating, %{field: new_value})
      {:ok, %EventRating{}}

      iex> update_event_rating(event_rating, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_event_rating(%EventRating{} = event_rating, attrs) do
    event_rating
    |> EventRating.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a event_rating.

  ## Examples

      iex> delete_event_rating(event_rating)
      {:ok, %EventRating{}}

      iex> delete_event_rating(event_rating)
      {:error, %Ecto.Changeset{}}

  """
  def delete_event_rating(%EventRating{} = event_rating) do
    Repo.delete(event_rating)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking event_rating changes.

  ## Examples

      iex> change_event_rating(event_rating)
      %Ecto.Changeset{data: %EventRating{}}

  """
  def change_event_rating(%EventRating{} = event_rating, attrs \\ %{}) do
    EventRating.changeset(event_rating, attrs)
  end

  alias Oplan.Event.EventComment

  @doc """
  Returns the list of event_comments.

  ## Examples

      iex> list_event_comments()
      [%EventComment{}, ...]

  """
  def list_event_comments do
    Repo.all(EventComment)
  end

  @doc """
  Gets a single event_comment.

  Raises `Ecto.NoResultsError` if the Event comment does not exist.

  ## Examples

      iex> get_event_comment!(123)
      %EventComment{}

      iex> get_event_comment!(456)
      ** (Ecto.NoResultsError)

  """
  def get_event_comment!(id), do: Repo.get!(EventComment, id)

  @doc """
  Creates a event_comment.

  ## Examples

      iex> create_event_comment(%{field: value})
      {:ok, %EventComment{}}

      iex> create_event_comment(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_event_comment(attrs \\ %{}) do
    %EventComment{}
    |> EventComment.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a event_comment.

  ## Examples

      iex> update_event_comment(event_comment, %{field: new_value})
      {:ok, %EventComment{}}

      iex> update_event_comment(event_comment, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_event_comment(%EventComment{} = event_comment, attrs) do
    event_comment
    |> EventComment.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a event_comment.

  ## Examples

      iex> delete_event_comment(event_comment)
      {:ok, %EventComment{}}

      iex> delete_event_comment(event_comment)
      {:error, %Ecto.Changeset{}}

  """
  def delete_event_comment(%EventComment{} = event_comment) do
    Repo.delete(event_comment)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking event_comment changes.

  ## Examples

      iex> change_event_comment(event_comment)
      %Ecto.Changeset{data: %EventComment{}}

  """
  def change_event_comment(%EventComment{} = event_comment, attrs \\ %{}) do
    EventComment.changeset(event_comment, attrs)
  end

  alias Oplan.Event.EventFollower

  @doc """
  Returns the list of event_followers.

  ## Examples

      iex> list_event_followers()
      [%EventFollower{}, ...]

  """
  def list_event_followers do
    Repo.all(EventFollower)
  end

  @doc """
  Gets a single event_follower.

  Raises `Ecto.NoResultsError` if the Event follower does not exist.

  ## Examples

      iex> get_event_follower!(123)
      %EventFollower{}

      iex> get_event_follower!(456)
      ** (Ecto.NoResultsError)

  """
  def get_event_follower!(id), do: Repo.get!(EventFollower, id)

  @doc """
  Creates a event_follower.

  ## Examples

      iex> create_event_follower(%{field: value})
      {:ok, %EventFollower{}}

      iex> create_event_follower(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_event_follower(attrs \\ %{}) do
    %EventFollower{}
    |> EventFollower.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a event_follower.

  ## Examples

      iex> update_event_follower(event_follower, %{field: new_value})
      {:ok, %EventFollower{}}

      iex> update_event_follower(event_follower, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_event_follower(%EventFollower{} = event_follower, attrs) do
    event_follower
    |> EventFollower.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a event_follower.

  ## Examples

      iex> delete_event_follower(event_follower)
      {:ok, %EventFollower{}}

      iex> delete_event_follower(event_follower)
      {:error, %Ecto.Changeset{}}

  """
  def delete_event_follower(%EventFollower{} = event_follower) do
    Repo.delete(event_follower)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking event_follower changes.

  ## Examples

      iex> change_event_follower(event_follower)
      %Ecto.Changeset{data: %EventFollower{}}

  """
  def change_event_follower(%EventFollower{} = event_follower, attrs \\ %{}) do
    EventFollower.changeset(event_follower, attrs)
  end
end
