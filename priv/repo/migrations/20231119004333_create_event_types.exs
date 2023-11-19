defmodule Oplan.Repo.Migrations.CreateEventTypes do
  use Ecto.Migration

  def change do
    create table(:event_types) do
      add :name, :string

      timestamps()
    end
  end
end
